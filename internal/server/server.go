package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/crpt/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	Engine *gin.Engine
	Config *config.Config
	Logger *zap.Logger
}

func NewServer(cfg *config.Config, logger *zap.Logger, handler *gin.Engine) *Server {
	return &Server{
		Engine: handler,
		Config: cfg,
		Logger: logger,
	}
}

func (s *Server) Start() error {
	serverAddr := s.Config.Server.Port
	if serverAddr == "" {
		serverAddr = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + serverAddr,
		Handler: s.Engine,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Sugar().Fatal("Failed to start server: ", zap.Error(err))
		}
	}()

	s.Logger.Sugar().Info("Server started", zap.String("address", serverAddr))

	<-quit
	s.Logger.Info("Server Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Sugar().Fatal("Server forced to shutdown: ", zap.Error(err))
	}

	s.Logger.Info("Server exited properly")
	return nil
}
