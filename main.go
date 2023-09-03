package main

import (
	"database/sql"
	"fmt"
	"log"
	"railway/database"
	"railway/model"
)

func main() {
	for {
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
			loadingStationsFromJsonFile()
		case "2":
			fmt.Println("Printing all stations from database")
			printAllStationsFromDatabase()
		case "3":
			fmt.Println("Removing all stations from database")
			removeAllStationFromDatabase()
		case "4":
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func removeAllStationFromDatabase() {
	var db *sql.DB
	var err error
	db, err = database.GetDBConnection()
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

func printAllStationsFromDatabase() {
	var db *sql.DB
	var err error
	var stations []model.Station

	db, err = database.GetDBConnection()

	stations, err = model.GetAllStations(db)
	if err != nil {
		log.Panicf(err.Error())
	}

	for _, station := range stations {
		station.Print()
	}
}

func loadingStationsFromJsonFile() {
	var stations []model.Station
	var err error
	stations, err = model.LoadStationsFromJsonFile("stations.json")
	if err != nil {
		log.Panicf(err.Error())
	}

	var db *sql.DB
	db, err = database.GetDBConnection()
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
