package main

import (
	"github.com/av-huette/avh-booking-system/config"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/logger"
	"github.com/av-huette/avh-booking-system/internal/models"
	"log/slog"
	"os"
)

type dbModels struct {
	account       *models.AccountModel
	accountOption *models.AccountOptionModel
	category      *models.CategoryModel
	unit          *models.UnitModel
	productGroup  *models.ProductGroupModel
	product       *models.ProductModel
}

type application struct {
	conf      *config.AppConfig
	log       *slog.Logger
	wsHandler *WebSocketHandler
	db        *database.DB
	dbModels  dbModels
}

func main() {
	wsHandler := NewWebSocketHandler(NewWebSocketService())

	dbPool, err := database.NewFromConfig()
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	app := &application{
		conf:      config.LoadConfig(),
		log:       logger.CreateLogger(),
		db:        dbPool,
		wsHandler: wsHandler,
		dbModels: dbModels{
			&models.AccountModel{DB: dbPool},
			&models.AccountOptionModel{DB: dbPool},
			&models.CategoryModel{DB: dbPool},
			&models.UnitModel{DB: dbPool},
			&models.ProductGroupModel{DB: dbPool},
			&models.ProductModel{DB: dbPool},
		},
	}

	err = app.serveHTTP()
	if err != nil {
		app.log.Error(err.Error())
		os.Exit(1)
	}
}
