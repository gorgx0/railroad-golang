package main

import (
	"database/sql"
	"log"
	"railway/database"
	"railway/model"
)

func main() {
	stations := model.LoadStations("stations.json")

	db := database.GetDBConnection()

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
