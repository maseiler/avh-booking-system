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
		Account:           &repo.AccountStore{DB: dbPool},
		Category:          &repo.CategoryStore{DB: dbPool},
		Location:          &repo.LocationStore{DB: dbPool},
		Product:           &repo.ProductStore{DB: dbPool},
		ProductGroup:      &repo.ProductGroupStore{DB: dbPool},
		ProductVisibility: &repo.ProductVisibilityStore{DB: dbPool},
		Unit:              &repo.UnitStore{DB: dbPool},
		Vat:               &repo.VatStore{DB: dbPool},
	}

	var handler slog.Handler
	if appConf.LogFormat == config.LogFormatJSON {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: appConf.LogLevel})
	} else {
		handler = tint.NewHandler(os.Stdout, &tint.Options{Level: appConf.LogLevel, TimeFormat: time.DateTime})
	}
	log := slog.New(handler)
	log.Debug("Logger initialised", slog.Any("level", appConf.LogLevel), slog.Any("format", appConf.LogFormat))

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
