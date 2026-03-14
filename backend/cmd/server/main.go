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
	if err := config.LoadEnv(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	appConf, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	dbConf, err := config.LoadDBConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	dbPool, err := database.New(dbConf.DBUser, dbConf.DBPassword, dbConf.DBHost, dbConf.DBPort, dbConf.DBName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
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
