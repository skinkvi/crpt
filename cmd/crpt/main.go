package main

import (
	"github.com/skinkvi/crpt/internal/config"
	database "github.com/skinkvi/crpt/pkg/db"
	"github.com/skinkvi/crpt/pkg/util/logger"
)

func main() {
	cfg, err := config.NewConfig("./config.yaml")
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}

	cfg.Logger.Info("Starting crpt...")

	db := database.Init()
	if db == nil {
		cfg.Logger.Fatal("Failed to connect to the database")
	}

	cfg.Logger.Info("Database connected")

	// TODO: init rabbitmq

	// TODO: start server

	// TODO: graceful shutdown
}
