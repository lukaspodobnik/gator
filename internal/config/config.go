package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filepath.Join(homeDir, configFileName))
	if err != nil {
		return Config{}, err
	}

	var cfg = Config{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	cfg.CurrentUserName = userName
	data, err := json.Marshal(*cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(homeDir, configFileName), data, os.ModePerm); err != nil {
		return err
	}

	return nil
}
