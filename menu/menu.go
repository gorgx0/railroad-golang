package menu

import (
	"database/sql"
	"fmt"
	"log"
	"railway/config"
	"railway/database"
	"railway/model"
)

func RemoveAllStationFromDatabase(dbConfig config.DatabaseConfig) error {
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

func PrintAllStationsFromDatabase(dbConfig config.DatabaseConfig) error {
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

func LoadingStationsFromJsonFile(config config.Config) error {
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
