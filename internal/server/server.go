// pkg/server/server.go
package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/skinkvi/crpt/internal/config"
	"github.com/skinkvi/crpt/pkg/util/logger"
)

func StartServer() {
	serverConfig := config.GetConfig().Server

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", serverConfig.Port),
		// Добавь другие настройки сервера по необходимости
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.GetLogger().Fatal(err.Error())
		}
	}()

	logger.GetLogger().Info("Server started")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.GetLogger().Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.GetLogger().Fatal(err.Error())
	}
	logger.GetLogger().Info("Server exited properly")
}
