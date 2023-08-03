package model

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Station struct {
	Lng  float64 `json:"lng"`
	Lat  float64 `json:"lat"`
	Name string  `json:"name"`
	Id   string  `json:"id"`
}

func (s Station) Print() {
	fmt.Printf("%s (%s)\n", s.Name, s.Id)
	println(s.Name, s.Id)
}

func LoadStations(filename string) []Station {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(file)
	var stations []Station
	err2 := json.Unmarshal(bytes, &stations)
	if err2 != nil {
		panic(err2)
	}
	return stations
}
