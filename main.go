package main

import (
	"database/sql"
	"fmt"
	"log"
	"railway/config"
	"railway/database"
	"railway/model"
)

func main() {
	currentConfig, err := config.ReadConfigFile("config.json")
	if err != nil {
		log.Panicf(err.Error())
	}

	for {
		menu(currentConfig)
	}
}

func menu(currentConfig config.Config) bool {
	fmt.Println("Menu:")
	fmt.Println("1. Load stations from JSON file into database")
	fmt.Println("2. Print all stations from database")
	fmt.Println("3. Remove all stations from database")
	fmt.Println("4. Exit")
	fmt.Println("Enter your choice: ")

	var choice string
	var err error
	_, err = fmt.Scanln(&choice)
	if err != nil {
		log.Panicf(err.Error())
	}

	switch choice {
	case "1":
		fmt.Println("Loading stations from JSON file into database")
		loadingStationsFromJsonFile(currentConfig.Database)
	case "2":
		fmt.Println("Printing all stations from database")
		printAllStationsFromDatabase(currentConfig.Database)
	case "3":
		fmt.Println("Removing all stations from database")
		removeAllStationFromDatabase(currentConfig.Database)
	case "4":
		fmt.Println("Exiting")
		return true
	default:
		fmt.Println("Invalid choice")
	}
	return false
}

func removeAllStationFromDatabase(dbConfig config.DatabaseConfig) {
	var db *sql.DB
	var err error
	db, err = database.GetDBConnection(dbConfig)
	if err != nil {
		log.Panicf(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panicf(err.Error())
		}
	}(db)

	_, err = db.Exec("DELETE FROM stations")
	if err != nil {
		log.Panicf(err.Error())
	}
}

func printAllStationsFromDatabase(dbConfig config.DatabaseConfig) {
	var db *sql.DB
	var err error
	var stations []model.Station

	db, err = database.GetDBConnection(dbConfig)

	stations, err = model.GetAllStations(db)
	if err != nil {
		log.Panicf(err.Error())
	}

	for _, station := range stations {
		station.Print()
	}
}

func loadingStationsFromJsonFile(dbConfig config.DatabaseConfig) {
	var stations []model.Station
	var err error
	stations, err = model.LoadStationsFromJsonFile("stations.json")
	if err != nil {
		log.Panicf(err.Error())
	}

	var db *sql.DB
	db, err = database.GetDBConnection(dbConfig)
	if err != nil {
		log.Panicf(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panicf(err.Error())
		}
	}(db)

	for _, station := range stations {
		err := station.Store(db)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
