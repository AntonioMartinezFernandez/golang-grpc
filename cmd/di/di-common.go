package di

import (
	"log/slog"

	"github.com/AntonioMartinezFernandez/golang-grpc/config"

	pkg_logger "github.com/AntonioMartinezFernandez/golang-grpc/pkg/logger"
)

type CommonServices struct {
	Config config.Config
	Logger *slog.Logger
}

func InitCommonServices() *CommonServices {
	config := initConfig()
	logger := pkg_logger.NewLogger(config)

	return &CommonServices{
		Config: config,
		Logger: logger,
	}
}

func initConfig() config.Config {
	return config.LoadEnvConfig()
}
