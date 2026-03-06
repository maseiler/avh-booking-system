package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/av-huette/avh-booking-system/internal/config"
	"github.com/av-huette/avh-booking-system/internal/database"
	"github.com/av-huette/avh-booking-system/internal/repo"
	"github.com/av-huette/avh-booking-system/internal/ws"
	"github.com/lmittmann/tint"
)

type application struct {
	conf      *config.AppConfig
	log       *slog.Logger
	db        *database.DB
	wsHandler *ws.Handler
}

func main() {
	config.LoadEnv()
	appConf := config.LoadConfig()
	dbConf := config.LoadDbConfig()

	dbPool, err := database.New(dbConf.DbUser, dbConf.DbPassword, dbConf.DbHost, dbConf.DbPort, dbConf.DbName)
	if err != nil {
		panic(err)
	}
	defer dbPool.Close()

	stores := ws.Stores{
		Account:           &repo.AccountModel{DB: dbPool},
		Category:          &repo.CategoryModel{DB: dbPool},
		Location:          &repo.LocationModel{DB: dbPool},
		Product:           &repo.ProductModel{DB: dbPool},
		ProductGroup:      &repo.ProductGroupModel{DB: dbPool},
		ProductVisibility: &repo.ProductVisibilityModel{DB: dbPool},
		Unit:              &repo.UnitModel{DB: dbPool},
		Vat:               &repo.VatModel{DB: dbPool},
	}

	logOpts := &tint.Options{Level: appConf.LogLevel, TimeFormat: time.DateTime}
	log := slog.New(tint.NewHandler(os.Stdout, logOpts))
	log.Debug(fmt.Sprintf("Log level: %s", logOpts.Level))

	app := &application{
		conf: appConf,
		log:  log,
		db:   dbPool,
	}

	app.wsHandler = ws.NewHandler(stores, app.log)

	if err := app.serveHTTP(); err != nil {
		app.log.Error(err.Error())
		os.Exit(1)
	}
}
