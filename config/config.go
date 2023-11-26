package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	AppServiceName string `env:"APP_SERVICE_NAME"`
	AppEnv         string `env:"APP_ENV"`
	AppVersion     string `env:"APP_VERSION"`

	LogLevel string `env:"LOG_LEVEL"`

	GrpcHost string `env:"GRPC_HOST"`
	GrpcPort int32  `env:"GRPC_PORT"`
}

func LoadEnvConfig() Config {
	var config Config

	godotenv.Load(".env")
	ctx := context.Background()
	if err := envconfig.Process(ctx, &config); err != nil {
		panic(err)
	}

	return config
}
