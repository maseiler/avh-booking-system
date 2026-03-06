package logger

import (
	"fmt"
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

func CreateLogger(level slog.Level) *slog.Logger {
	options := &tint.Options{
		Level:      level,
		TimeFormat: time.DateTime,
	}

	log := slog.New(tint.NewHandler(os.Stdout, options))
	log.Debug(fmt.Sprintf("Log level: %s", options.Level))

	return log
}
