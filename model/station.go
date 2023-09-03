package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Station struct {
	Id       int16 `json:"id"`
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Name string `json:"name"`
}

func (s Station) Print() {
	fmt.Printf("%s (%s)\n", s.Name, s.Id)
	println(s.Name, s.Id)
}

func (s Station) Store(db *sql.DB) {
	_, err := db.Exec("INSERT INTO stations (id, name, lat, lng) VALUES ($1, $2, $3, $4)", s.Id, s.Name, s.Location.Latitude, s.Location.Longitude)
	if err != nil {
		log.Panicf(err.Error())
	}
}

func LoadStations(filename string) []Station {
	file, err := os.Open(filename)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panic(err)
		}
	}(file)

	if err != nil {
		log.Panic(err)
	}
	bytes, err := io.ReadAll(file)
	var stations []Station
	err2 := json.Unmarshal(bytes, &stations)
	if err2 != nil {
		log.Panic(err2)
	}
	return stations
}

func GetAllStations(db *sql.DB) []Station {
	rows, err := db.Query("SELECT id, name, lat, lng FROM stations")
	if err != nil {
		log.Panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic(err)
		}
	}(rows)
	var stations []Station
	for rows.Next() {
		var station Station
		err := rows.Scan(&station.Id, &station.Name, &station.Location.Latitude, &station.Location.Longitude)
		if err != nil {
			log.Panic(err)
		}
		stations = append(stations, station)
	}
	return stations
}
