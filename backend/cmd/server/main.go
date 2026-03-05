package main

import (
	"github.com/av-huette/avh-booking-system/config"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/logger"
	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/ws"
	"log/slog"
	"os"
)

type application struct {
	conf      *config.AppConfig
	log       *slog.Logger
	db        *database.DB
	WsHandler *WebSocketHandler
}

func main() {
	dbPool, err := database.NewFromConfig()
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	stores := ws.Stores{
		Account:           &models.AccountModel{DB: dbPool},
		Category:          &models.CategoryModel{DB: dbPool},
		Location:          &models.LocationModel{DB: dbPool},
		Product:           &models.ProductModel{DB: dbPool},
		ProductGroup:      &models.ProductGroupModel{DB: dbPool},
		ProductVisibility: &models.ProductVisibilityModel{DB: dbPool},
		Unit:              &models.UnitModel{DB: dbPool},
		Vat:               &models.VatModel{DB: dbPool},
	}

	app := &application{
		conf: config.LoadConfig(),
		log:  logger.CreateLogger(),
		db:   dbPool,
	}

	app.WsHandler = NewWebSocketHandler(app, stores)

	if err := app.serveHTTP(); err != nil {
		app.log.Error(err.Error())
		os.Exit(1)
	}
}
