package main

import (
	"context"
	"fmt"
	"os"

	"github.com/av-huette/avh-booking-system/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const migrationsDir = "./db/migrations"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: migrate <command> [args]")
		fmt.Fprintln(os.Stderr, "commands: up, down, status, version, redo, reset")
		os.Exit(1)
	}

	if err := config.LoadEnv(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	dbConf, err := config.LoadDBConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		dbConf.DBHost, dbConf.DBPort, dbConf.DBName, dbConf.DBUser, dbConf.DBPassword)

	connConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	db := stdlib.OpenDB(*connConfig)
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]
	if err := goose.RunContext(context.Background(), command, db, migrationsDir, args...); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
