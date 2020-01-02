package config

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	Configs "github.com/timechanger/config"
)

var configs Configs.Configuration

func main() {

	configs, _ = Configs.PrepareConfigs()
	initTimer(configs.Interval, configs.TimeAddition)

}

func initTimer(interval uint64, timeAddition uint64) {

	//go func() {
	for now := range time.Tick(time.Second) {
		newTime := time.Now().Add(time.Second * time.Duration(timeAddition))
		fmt.Println(now, ": Setting system date to ", newTime)
		err := SetSystemDateUsingExec(newTime)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
	}
	//}()

}

func disableAutoSync() {

	f, errOpen := os.OpenFile("/etc/default/ntp", os.O_APPEND|os.O_WRONLY, 0644)
	if errOpen != nil {
		panic(errOpen)
	}

	_, errWrite := f.WriteString("exit")
	if errWrite != nil {
		panic(errWrite)
	}

	f.Close()

}

func SetSystemDateUsingExec(newTime time.Time) error {

	_, lookErr := exec.LookPath("date")
	if lookErr != nil {
		return lookErr
	}
	dateString := newTime.Format("2006-01-2 15:4:5")
	dateCmd := "date --set=\"" + dateString + "\""
	cmd := exec.Command("/bin/sh", "-c", "sudo "+dateCmd)
	return cmd.Run()
}

func SetSystemDateUsingSysCall(newTime time.Time) error {

	binary, lookErr := exec.LookPath("date")
	//binary, lookErr := exec.LookPath("timedatectl")
	if lookErr != nil {
		fmt.Printf("Date binary not found, cannot set system date: %s\n", lookErr.Error())
		return lookErr
	}
	dateString := newTime.Format("2006/01/02 15:04:05")
	args := []string{"--set=\"" + dateString + "\""}
	env := os.Environ()
	return syscall.Exec(binary, args, env)

}
