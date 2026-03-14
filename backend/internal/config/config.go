package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	DefaultHTTPPort = 8081
)

type AppConfig struct {
	LogLevel     slog.Level
	HTTPPort     int
	FrontendPath string
}

type DBConfig struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file: " + err.Error())
	}
}

func LoadEnvFromFile(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
}

func LoadConfig() *AppConfig {
	conf := &AppConfig{}
	conf.LogLevel = getLogLevel("AVHBS_LOG_LEVEL", slog.LevelInfo)
	conf.HTTPPort = getInt("AVHBS_HTTP_PORT", DefaultHTTPPort)
	conf.FrontendPath = getString("AVHBS_FRONTEND_PATH", "")

	return conf
}

func LoadDBConfig() (*DBConfig, error) {
	conf := &DBConfig{
		DBHost:     getString("DB_HOST", ""),
		DBPort:     getInt("DB_PORT", 0),
		DBName:     getString("DB_NAME", ""),
		DBUser:     getString("DB_USER", ""),
		DBPassword: getString("DB_PASSWORD", ""),
	}

	var missing []string
	if conf.DBHost == "" {
		missing = append(missing, "DB_HOST")
	}
	if conf.DBPort == 0 {
		missing = append(missing, "DB_PORT")
	}
	if conf.DBName == "" {
		missing = append(missing, "DB_NAME")
	}
	if conf.DBUser == "" {
		missing = append(missing, "DB_USER")
	}
	if conf.DBPassword == "" {
		missing = append(missing, "DB_PASSWORD")
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required env vars: %s", strings.Join(missing, ", "))
	}

	return conf, nil
}

func getString(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func getInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

func getLogLevel(key string, defaultValue slog.Level) slog.Level {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	m := make(map[string]slog.Level)
	m["DEBUG"] = slog.LevelDebug
	m["INFO"] = slog.LevelInfo
	m["WARN"] = slog.LevelWarn
	m["ERROR"] = slog.LevelError

	logLevel, exists := m[value]
	if exists {
		return logLevel
	}

	return defaultValue
}
