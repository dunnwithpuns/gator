package config

import (
	"os"
	"encoding/json"
)

func READ() (Config, error) {

	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}
	
	config := Config{}
	if err := json.Unmarshal(data, &config); err != nil{
		return Config{}, err
	}

	return config, nil
}
