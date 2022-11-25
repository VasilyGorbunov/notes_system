package main

import (
	"github.com/VasilyGorbunov/notes_system/api_service/internal/config"
	"github.com/VasilyGorbunov/notes_system/api_service/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	logger.Println("config initializing")
	cfg := config.GetConfig()
}
