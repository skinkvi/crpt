package main

import (
	"github.com/skinkvi/crpt/internal/config"
	"github.com/skinkvi/crpt/internal/handlers"
	"github.com/skinkvi/crpt/internal/server"
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

	// rabbitmqConn, err := rabbitmq.InitRabbitMQ()
	// if err != nil {
	// 	cfg.Logger.Fatal("Failed to connect to RabbitMQ")
	// }
	// defer rabbitmqConn.Close()

	handlers.InitHandlers()

	cfg.Logger.Info("Handlers initialized")

	server.StartServer()

	// Блокирующий вызов для ожидания завершения сервера
	select {}
}

