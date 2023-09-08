package main

import (
	"log"
	"railway/config"
	"railway/menu"
)

func main() {
	currentConfig, err := config.ReadConfigFromFile("config.json")
	if err != nil {
		log.Panicf(err.Error())
	}
	for {
		err := menu.Menu(currentConfig)
		if err != nil {
			log.Panicf(err.Error())
		}
	}
}
