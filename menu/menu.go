package menu

import (
	"database/sql"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"log"
	"os"
	"railway/config"
	"railway/database"
	"railway/model"
)

func Menu(currentConfig config.Config) error {
	fmt.Println("Menu:")
	fmt.Println("1. Load stations from JSON file into database")
	fmt.Println("2. Print all stations from database")
	fmt.Println("3. Remove all stations from database")
	fmt.Println("4. Display map")
	fmt.Println("X. Exit")
	fmt.Println("Enter your choice: ")

	var choice string
	var err error
	_, err = fmt.Scanln(&choice)
	if err != nil {
		return err
	}

	switch choice {
	case "1":
		fmt.Println("Loading stations from JSON file into database")
		return loadingStationsFromJsonFile(currentConfig)
	case "2":
		fmt.Println("Printing all stations from database")
		return printAllStationsFromDatabase(currentConfig.Database)
	case "3":
		fmt.Println("Removing all stations from database")
		return removeAllStationFromDatabase(currentConfig.Database)
	case "4":
		fmt.Println("Showing map")
		image, err := model.GetMapImage(currentConfig.Database)
		if err != nil {
			return err
		}
		app := app.New()
		w := app.NewWindow("Railway")

		mapCanvas := canvas.NewImageFromImage(image)
		mapCanvas.FillMode = canvas.ImageFillOriginal
		w.SetContent(mapCanvas)
		w.ShowAndRun()

		return nil
	case "x", "X":
		fmt.Println("Exiting")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
	return nil
}

func removeAllStationFromDatabase(dbConfig config.DatabaseConfig) error {
	var (
		db  *sql.DB
		err error
	)
	db, err = database.GetDBConnection(dbConfig)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panicf(err.Error())
		}
	}(db)

	_, err = db.Exec("DELETE FROM stations WHERE true")
	if err != nil {
		return err
	}
	return nil
}

func printAllStationsFromDatabase(dbConfig config.DatabaseConfig) error {
	var db *sql.DB
	var err error
	var stations []model.Station

	db, err = database.GetDBConnection(dbConfig)
	if err != nil {
		return err
	}

	stations, err = model.GetAllStations(db)
	if err != nil {
		return err
	}

	fmt.Println("======== Stations Start ========")
	for _, station := range stations {
		station.Print()
	}
	fmt.Println("======== Stations End ========")
	return nil
}

func loadingStationsFromJsonFile(config config.Config) error {
	var stations []model.Station
	var err error
	stations, err = model.LoadStationsFromJsonFile(config.StationsFile)
	if err != nil {
		return err
	}

	var db *sql.DB
	db, err = database.GetDBConnection(config.Database)
	if err != nil {
		return err
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
			return err
		}
	}
	return nil
}
