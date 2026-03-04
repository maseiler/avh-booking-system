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
	SchemasPath  string
}

type DbConfig struct {
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
}

func LoadConfig() *AppConfig {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	conf := &AppConfig{}
	conf.LogLevel = getLogLevel("AVHBS_LOG_LEVEL", slog.LevelInfo)
	conf.HttpPort = getInt("AVHBS_HTTP_PORT", DefaultHTTPPort)
	conf.FrontendPath = getString("AVHBS_FRONTEND_PATH", "")
	conf.SchemasPath = getString("AVHBS_SCHEMAS_PATH", "../schemas")

	return conf
}

func LoadDbConfigFromRootEnv() *DbConfig {
	// load environment variables from root .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	return createDbConfig()
}

func LoadDbConfigFromFileEnv(filePath string) *DbConfig {
	// load environment variables from absolute path to .env file
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

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
