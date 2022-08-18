package main

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	avhbsConfig := getConfig()
	// Create an instance of the app structure
	app := NewApp(avhbsConfig)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "AVH Booking System",
		Width:  1024,
		Height: 768,
		// MinWidth:          720,
		// MinHeight:         570,
		// MaxWidth:          1280,
		// MaxHeight:         740,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnShutdown:        app.shutdown,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func getConfig() AvhbsConfig {
	file, err := os.Open("/usr/share/avhbs/config")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal("Can not find AVHBS config file.")
		}
		log.Fatal(err)
	}
	defer file.Close()

	url := ""
	clientName := ""
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if i == 0 {
			url = scanner.Text()
		}
		if i == 1 {
			clientName = scanner.Text()
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	avhbsConfig := AvhbsConfig{url, clientName}

	fmt.Println(avhbsConfig)

	return avhbsConfig
	// TODO: check if client name in DB
}
