package config

import (
	"os"
	"encoding/json"
)

func READ() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(home + "/.gatorconfig.json")
	if err != nil {
		return Config{}, err
	}
	
	config := Config{}
	if err := json.Unmarshal(data, &config); err != nil{
		return Config{}, err
	}

	return config, nil
}
