package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func readConfigFile(filename string) (Config, error) {
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

func main() {
	config, err := readConfigFile("config.json")
	if err != nil {
		fmt.Println("Error reading configuration:", err)
		return
	}

	fmt.Println("Database Host:", config.Database.Host)
	fmt.Println("Database Port:", config.Database.Port)
	fmt.Println("Database User:", config.Database.User)
	fmt.Println("Database Name:", config.Database.Db)
}
