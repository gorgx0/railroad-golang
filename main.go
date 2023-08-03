package main

import (
	"log"
	"railway/database"
	"railway/model"
)

func main() {
	stations := model.LoadStations("stations.json")
	db := database.GetDBConnection()
	for _, station := range stations {
		_, err := db.Exec("INSERT INTO stations (id, name, lat, lng) VALUES ($1, $2, $3, $4)", station.Id, station.Name, station.Lat, station.Lng)
		if err != nil {
			log.Panicf(err.Error())
		}
	}
}
