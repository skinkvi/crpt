package app

import (
	"github.com/gin-gonic/gin"
	"github.com/skinkvi/crpt/internal/config"
	"github.com/skinkvi/crpt/internal/handlers"
	"github.com/skinkvi/crpt/internal/server"
	"github.com/skinkvi/crpt/internal/services"
	"github.com/skinkvi/crpt/pkg/storage"
	"go.uber.org/zap"
)

type App struct {
	Config  *config.Config
	Logger  *zap.Logger
	Server  *server.Server
	Router  *gin.Engine
	Handler *handlers.Handler
	Service *services.Service
	Storage storage.Storage
}

func NewApp(configPath string, logger *zap.Logger) (*App, error) {
	cfg, err := config.NewConfig(configPath, logger)
	if err != nil {
		logger.Error("Failed to load config", zap.Error(err))
		return nil, err
	}

	// Инициализация зависимостей
	dbStorage := storage.NewStorage(cfg, logger)
	service := services.NewService(dbStorage, logger)
	handler := handlers.NewHandler(*service, logger)

	// Инициализация Gin
	router := gin.Default()

	// Регистрация маршрутов
	router.GET("/", handler.Home)
	router.GET("/crypto", handler.Crypto)

	// Создание сервера
	srv := server.NewServer(cfg, logger, router)

	app := &App{
		Config:  cfg,
		Logger:  logger,
		Server:  srv,
		Router:  router,
		Handler: handler,
		Service: service,
		Storage: dbStorage,
	}

	return app, nil
}

func (a *App) Run() error {
	return a.Server.Start()
}
