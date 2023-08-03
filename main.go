package main

import "railway/model"

func main() {
	stations := model.LoadStations("stations.json")
	for _, station := range stations {
		station.Print()
	}
}
