package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"log"
	"railway/config"
	"railway/menu"
)

func main() {
	currentConfig, err := config.ReadConfigFromFile("config.json")
	if err != nil {
		log.Panicf(err.Error())
	}

	app := app.New()
	w := app.NewWindow("Railway")
	w.SetContent(widget.NewLabel("Hello Fyne!"))
	w.ShowAndRun()
	for {
		err := menu.Menu(currentConfig)
		if err != nil {
			log.Panicf(err.Error())
		}
	}
}
