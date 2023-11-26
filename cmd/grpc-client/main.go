package main

import (
	"github.com/AntonioMartinezFernandez/golang-grpc/cmd/di"
)

func main() {
	commonDi := di.InitCommonServices()

	commonDi.Logger.Info("Starting gRPC client...")
}
