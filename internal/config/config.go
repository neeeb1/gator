package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_string"`
}

func getConfigFilePath() (string, error) {
	var filepath string
	home, err := os.UserHomeDir()
	if err != nil {
		return filepath, err
	}

	filepath = fmt.Sprintf("%s/%s", home, configFileName)
	return filepath, nil
}

func Read() (Config, error) {
	var cfg Config

	filepath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	os.WriteFile(filepath, data, 0644)
	return nil
}

func (cfg Config) SetUser(user string) {
	cfg.CurrentUser = user
	write(cfg)
}
