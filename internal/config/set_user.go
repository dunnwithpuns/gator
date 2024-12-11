package config

import (
	"encoding/json"
	"os"
)

func SetUser(userName string) error {

	cfg := Config {
		Current_User_Name: userName,
	}

	err := write(cfg)
	if err != nil {
		return err
	}

	return nil
}

func write(cfg Config) error {

	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	os.WriteFile(configPath, jsonData, 0644)

	return nil
}

