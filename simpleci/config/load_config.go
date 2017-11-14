package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadConfig load the config from jsonConfigPath
func LoadConfig(configFilePath string) (Config, error) {
	var config Config
	configFile, err := os.Open(configFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}
