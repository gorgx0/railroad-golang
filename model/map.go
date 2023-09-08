package model

import (
	"fmt"
	sm "github.com/flopp/go-staticmaps"
	"github.com/golang/geo/s2"
	"golang.org/x/image/colornames"
	"image/jpeg"
	"log"
	"os"
	"railway/config"
	"railway/database"
)

type Rectangle struct {
	MinLat float64
	MaxLat float64
	MinLng float64
	MaxLng float64
}

func GetRectangle(dbConfig config.DatabaseConfig) (Rectangle, error) {
	rectangle := Rectangle{
		MinLat: 0,
		MaxLat: 0,
		MinLng: 0,
		MaxLng: 0,
	}
	db, err := database.GetDBConnection(dbConfig)
	if err != nil {
		return rectangle, err
	}
	stations, err := GetAllStations(db)
	if err != nil {
		return rectangle, err
	}
	rectangle.MinLat = stations[0].Location.Latitude
	rectangle.MaxLat = stations[0].Location.Latitude
	rectangle.MinLng = stations[0].Location.Longitude
	rectangle.MaxLng = stations[0].Location.Longitude

	for _, station := range stations {
		if station.Location.Latitude < rectangle.MinLat {
			rectangle.MinLat = station.Location.Latitude
		}
		if station.Location.Latitude > rectangle.MaxLat {
			rectangle.MaxLat = station.Location.Latitude
		}
		if station.Location.Longitude < rectangle.MinLng {
			rectangle.MinLng = station.Location.Longitude
		}
		if station.Location.Longitude > rectangle.MaxLng {
			rectangle.MaxLng = station.Location.Longitude
		}
	}
	ctx := sm.NewContext()
	ctx.SetSize(800, 600)
	bbox, err := sm.CreateBBox(rectangle.MinLat, rectangle.MinLng, rectangle.MaxLat, rectangle.MaxLng)
	if err != nil {
		log.Panic(err)
	}
	ctx.SetBoundingBox(*bbox)

	for _, station := range stations {
		pos := s2.LatLngFromDegrees(station.Location.Latitude, station.Location.Longitude)
		marker := sm.NewMarker(pos, colornames.Red, 10)
		marker.SetLabelColor(colornames.Black)
		marker.Label = fmt.Sprintf("%s (%d)", station.Name, station.Id)
		marker.LabelYOffset = 2
		ctx.AddObject(marker)
	}

	img, err := ctx.Render()
	if err != nil {
		log.Panic(err)
	}
	out, err := os.Create("map.png")
	if err != nil {
		log.Panic(err)
	}

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		log.Panic(err)
	}
	log.Default().Println("Map saved to map.png")
	return rectangle, nil
}

func (r Rectangle) Print() {
	fmt.Printf("MinLat: %f MaxLat: %f MinLng: %f MaxLng: %f\n", r.MinLat, r.MaxLat, r.MinLng, r.MaxLng)
}
