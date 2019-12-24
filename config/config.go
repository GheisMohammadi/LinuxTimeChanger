package config

import (
	"encoding/json"
	"os"
)

//Configuration  returns configs
type Configuration struct {
	StartDate    string
	Interval     uint64
	TimeAddition uint64
}

//PrepareConfigs prepares config from config.json
func PrepareConfigs(configFile string) (Configuration, error) {

	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, nil

}
