package config

import "os"

type Config struct {
	db_URL string `json:"db_url"`
	Current_User_Name string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + configFileName, nil
}
