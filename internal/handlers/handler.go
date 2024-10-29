package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skinkvi/crpt/internal/services"
	"go.uber.org/zap"
)

type Handler struct {
	Service *services.Service
	Logger  *zap.Logger
}

func NewHandler(service services.Service, logger *zap.Logger) *Handler {
	return &Handler{
		Service: &service,
		Logger:  logger,
	}
}

func (h *Handler) Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}

func (h *Handler) Crypto(c *gin.Context) {
	cryptoName := c.Query("name")
	if cryptoName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing crypto name",
		})
		h.Logger.Error("Not found crypto name")
		return
	}

	cryptoData, err := h.Service.GetCryptoData(cryptoName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("Falied to get data from crypto name")
		return
	}

	h.Logger.Sugar().Info("Successful found crypto data: ", cryptoData)
	c.JSON(http.StatusOK, cryptoData)
}
