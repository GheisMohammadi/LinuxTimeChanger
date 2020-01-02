package config

import (
	"encoding/json"
	"os"
)

//Configuration  returns configs
type Configuration struct {
	Interval     uint64
	TimeAddition uint64
}

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
