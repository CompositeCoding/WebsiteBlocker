package main

import (
	"embed"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	log.SetOutput(file)
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:       "Website Blocker",
		MinWidth:    900,
		Width:       1024,
		Height:      768,
		StartHidden: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 18, G: 18, B: 18, A: 1},
		OnStartup:        app.startup,
		Windows: &windows.Options{
			WindowIsTranslucent:  true,
			WebviewGpuIsDisabled: true,
			DisableWindowIcon:    true,
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
