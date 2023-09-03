package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database     DatabaseConfig `json:"database"`
	StationsFile string         `json:"stations_file"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func ReadConfigFromFile(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}

	return config, nil
}
