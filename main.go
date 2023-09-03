package main

import (
	"database/sql"
	"log"
	"railway/database"
	"railway/model"
)

func main() {
	var stations, err = model.LoadStations("stations.json")
	if err != nil {
		log.Panicf(err.Error())
	}

	db, err := database.GetDBConnection()
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
		station.Store(db)
	}
}
