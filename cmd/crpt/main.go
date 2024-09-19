package main

import (
	"github.com/skinkvi/crpt/internal/config"
	database "github.com/skinkvi/crpt/pkg/db"
	"github.com/skinkvi/crpt/pkg/util/logger"
)

func main() {
	config.NewConfig("config.yaml")
	logger.GetLogger().Info("Starting crpt...")

	database.Init()
	logger.GetLogger().Info("Database connected")

	// TODO: init rabbitmq

	// TODO: start server

	// TODO: graceful shutdown
}
