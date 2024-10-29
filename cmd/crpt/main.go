package main

import (
	"github.com/skinkvi/crpt/internal/app"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	application, err := app.NewApp("./config.yaml", logger)
	if err != nil {
		logger.Sugar().Fatal("Failed to initialize app", zap.Error(err))
	}

	if err := application.Run(); err != nil {
		logger.Sugar().Fatal("Application exited with error", zap.Error(err))
	}
}
