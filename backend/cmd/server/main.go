package main

import (
	"github.com/av-huette/avh-booking-system/config"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/logger"
	"github.com/av-huette/avh-booking-system/internal/models"
	"log/slog"
	"os"
)

type application struct {
	conf      *config.AppConfig
	log       *slog.Logger
	db        *database.DB
	DbModels  *models.DbModels
	WsHandler *WebSocketHandler
}

func main() {
	dbPool, err := database.NewFromConfig()
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	app := &application{
		conf: config.LoadConfig(),
		log:  logger.CreateLogger(),
		db:   dbPool,
		DbModels: &models.DbModels{
			Account:       models.AccountModel{DB: dbPool},
			AccountOption: models.AccountOptionModel{DB: dbPool},
			Category:      models.CategoryModel{DB: dbPool},
			Unit:          models.UnitModel{DB: dbPool},
			ProductGroup:  models.ProductGroupModel{DB: dbPool},
			Product:       models.ProductModel{DB: dbPool},
		},
	}

	app.WsHandler = NewWebSocketHandler(app)

	if err := app.serveHTTP(); err != nil {
		app.log.Error(err.Error())
		os.Exit(1)
	}
}
