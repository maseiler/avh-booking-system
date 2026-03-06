package config

import (
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"strconv"
)

const (
	DefaultHTTPPort = 8081
)

type AppConfig struct {
	LogLevel     slog.Level
	HttpPort     int
	FrontendPath string
}

type DbConfig struct {
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
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
	conf.HttpPort = getInt("AVHBS_HTTP_PORT", DefaultHTTPPort)
	conf.FrontendPath = getString("AVHBS_FRONTEND_PATH", "")

	return conf
}

func LoadDbConfig() *DbConfig {
	return createDbConfig()
}

func createDbConfig() *DbConfig {
	conf := &DbConfig{}
	conf.DbHost = getString("DB_HOST", "INVALID")
	conf.DbPort = getInt("DB_PORT", 0)
	conf.DbName = getString("DB_NAME", "INVALID")
	conf.DbUser = getString("DB_USER", "INVALID")
	conf.DbPassword = getString("DB_PASSWORD", "INVALID")

	return conf
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

func getBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return boolValue
}

func getLogLevel(key string, defaultValue slog.Level) slog.Level {
	value, exists := os.LookupEnv(key)
	if !exists {
		return slog.LevelInfo
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
