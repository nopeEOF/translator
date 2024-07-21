package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Lang string `json:"lang"`
	Url  string `json:"url"`
}

func getConfig(configPath string) (Config, error) {
	var config Config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	json.Unmarshal(data, &config)
	return config, nil
}
