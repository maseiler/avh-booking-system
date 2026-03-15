package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	DefaultHTTPPort = 8081
)

type LogFormat int

const (
	LogFormatText LogFormat = iota
	LogFormatJSON
)

type AppConfig struct {
	LogLevel     slog.Level
	LogFormat    LogFormat
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

func LoadEnv() error {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("loading .env file: %w", err)
	}
	return nil
}

func LoadEnvFromFile(path string) error {
	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("loading .env file %q: %w", path, err)
	}
	return nil
}

func LoadConfig() (*AppConfig, error) {
	httpPort, err := getInt("AVHBS_HTTP_PORT", DefaultHTTPPort)
	if err != nil {
		return nil, fmt.Errorf("AVHBS_HTTP_PORT: %w", err)
	}

	conf := &AppConfig{
		LogLevel:     getLogLevel("AVHBS_LOG_LEVEL", slog.LevelInfo),
		LogFormat:    getLogFormat("AVHBS_LOG_FORMAT", LogFormatText),
		HTTPPort:     httpPort,
		FrontendPath: getString("AVHBS_FRONTEND_PATH", ""),
	}

	return conf, nil
}

func LoadDBConfig() (*DBConfig, error) {
	dbPort, err := getInt("DB_PORT", 0)
	if err != nil {
		return nil, fmt.Errorf("DB_PORT: %w", err)
	}

	conf := &DBConfig{
		DBHost:     getString("DB_HOST", ""),
		DBPort:     dbPort,
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

func getInt(key string, defaultValue int) (int, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue, nil
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value %q: %w", value, err)
	}

	return intValue, nil
}

func getLogFormat(key string, defaultValue LogFormat) LogFormat {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	switch strings.ToUpper(value) {
	case "JSON":
		return LogFormatJSON
	case "TEXT":
		return LogFormatText
	default:
		return defaultValue
	}
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
