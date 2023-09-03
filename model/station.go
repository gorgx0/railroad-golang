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
	fmt.Printf("%s (%d)\n", s.Name, s.Id)
}

func (s Station) Store(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO stations (id, name, lat, lng) VALUES ($1, $2, $3, $4)", s.Id, s.Name, s.Location.Latitude, s.Location.Longitude)
	return err
}

func LoadStationsFromJsonFile(filename string) ([]Station, error) {
	var file, err = os.Open(filename)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	var stations []Station
	err = json.Unmarshal(bytes, &stations)
	if err != nil {
		return nil, err
	}
	return stations, nil
}

func GetAllStations(db *sql.DB) ([]Station, error) {
	rows, err := db.Query("SELECT id, name, lat, lng FROM stations")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var stations []Station
	for rows.Next() {
		var station Station
		err := rows.Scan(&station.Id, &station.Name, &station.Location.Latitude, &station.Location.Longitude)
		if err != nil {
			return stations, err
		}
		stations = append(stations, station)
	}
	return stations, nil
}
