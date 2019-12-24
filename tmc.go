package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
	Configs "timechanger/config"

	ntp "github.com/beevik/ntp"
)

var configs Configs.Configuration

var ntpTime time.Time

func main() {

	configs, _ = Configs.PrepareConfigs("config.json")

	//Setup ntp
	var errSetupNTP error
	ntpTime, errSetupNTP = ntp.Time("time.apple.com")
	if errSetupNTP != nil {
		fmt.Println("setup ntp error: ", errSetupNTP)
	}
	fmt.Println("Current real local time: ", ntpTime.Local())
	fmt.Println("Current real UTC time: ", ntpTime.UTC())
	//diff := ntpTime.Local().Sub(ntpTime.UTC())
	//println("Time diff (local-UTC) is: ", diff.Round)

	//Set Start Date
	layout := "2006-01-02 15:04:05"
	startDateStr := configs.StartDate
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		fmt.Println("Start time parse error:", err)
	}
	//startDate = startDate.Add(diff)
	fmt.Println("Start date: ", startDate)

	//Set system time on start date
	SetSystemDateUsingExec(startDate)

	//startDate := new time.Date(configs.StartDate)
	initTimer(configs.Interval, configs.TimeAddition)

}

func initTimer(interval uint64, timeAddition uint64) {

	//go func() {
	sec := uint64(0)

	for now := range time.Tick(time.Second) {
		sec++
		if sec >= interval {
			sec = 0
			newTime := time.Now().Add(time.Second * time.Duration(timeAddition))
			fmt.Println(now, ": Setting system date to ", newTime)
			err := SetSystemDateUsingExec(newTime)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}

			currentReadTime := ntpTime.UTC()
			if newTime.Unix() > currentReadTime.Unix() {
				println("Reached current time!")
				break
			}
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
