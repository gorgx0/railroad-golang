package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"log"
	"railway/config"
	"railway/menu"
	"railway/model"
)

func main() {
	currentConfig, err := config.ReadConfigFromFile("config.json")
	if err != nil {
		log.Panicf(err.Error())
	}

	app := app.New()
	window := app.NewWindow("Railway")
	window.Resize(fyne.NewSize(800, 600))
	window.CenterOnScreen()
	window.SetContent(widget.NewLabel("Hello World"))
	window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Load stations", func() {
				err := menu.LoadingStationsFromJsonFile(currentConfig)
				if err != nil {
					log.Panicf(err.Error())
				}
			}),
			fyne.NewMenuItem("Print stations", func() {
				err := menu.PrintAllStationsFromDatabase(currentConfig.Database)
				if err != nil {
					log.Panicf(err.Error())
				}
			}),
			fyne.NewMenuItem("Remove stations", func() {
				err := menu.RemoveAllStationFromDatabase(currentConfig.Database)
				if err != nil {
					log.Panicf(err.Error())
				}
			}),
			fyne.NewMenuItem("Show map", func() {
				image, err := model.GetMapImage(currentConfig.Database)
				if err != nil {
					log.Panicf(err.Error())
				}
				mapCanvas := canvas.NewImageFromImage(image)
				mapCanvas.FillMode = canvas.ImageFillOriginal
				window.SetContent(mapCanvas)
			}),
			fyne.NewMenuItem("Quit", func() {
				app.Quit()
			}))))
	window.ShowAndRun()

	for {
		err := menu.Menu(currentConfig)
		if err != nil {
			log.Panicf(err.Error())
		}
	}
}
