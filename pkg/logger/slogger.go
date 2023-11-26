package logger

import (
	"os"

	"log/slog"

	"github.com/AntonioMartinezFernandez/golang-grpc/config"
)

func NewLogger(cfg config.Config) *slog.Logger {
	var logLevel slog.Level
	switch level := cfg.LogLevel; level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelDebug
	}

	jsonHandler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: logLevel,
		},
	)

	return slog.New(jsonHandler)
}

func NewNullLogger() *slog.Logger {
	return &slog.Logger{}
}
