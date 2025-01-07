package config

import (
	"fmt"
	"os"
)

type Config struct {
	DB_URL string `json:"db_url"`
	Current_User_Name string `json:"current_user_name"`
}

const databaseURL = "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
const configFileName = "/.gatorconfig.json"

func getConfigFilePath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error locating home directory: %v", err)
	}

	return home + configFileName, nil
}
