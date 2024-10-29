package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/skinkvi/crpt/internal/models"
	"github.com/skinkvi/crpt/pkg/storage"
	"go.uber.org/zap"
)

type Service struct {
	Storage storage.Storage
	Logger  *zap.Logger
}

func NewService(storage storage.Storage, logger *zap.Logger) *Service {
	return &Service{
		Storage: storage,
		Logger:  logger,
	}
}

func (s *Service) GetCryptoData(cryptoName string) (models.CryptoData, error) {
	data, err := s.Storage.GetCryptoData(cryptoName)
	if err != nil {
		s.Logger.Sugar().Error("Data not found in DB of error occurred: ", err)
		return s.FecthAndSaveCryptoData(cryptoName)
	}

	if time.Since(data.UpdatedAt) > 60*time.Second {
		s.Logger.Info("Data is outdate, fetching new data")
		return s.FecthAndSaveCryptoData(cryptoName)
	}

	s.Logger.Info("Returning data from database")
	return data, nil
}

func (s *Service) FecthAndSaveCryptoData(cryptoName string) (models.CryptoData, error) {
	data, err := s.FecthCryptoDataFromAPI(cryptoName)
	if err != nil {
		return models.CryptoData{}, err
	}

	err = s.Storage.SaveCryptoData(data)
	if err != nil {
		s.Logger.Sugar().Error("Falied to save data to database: ", err)
		return models.CryptoData{}, err
	}

	return data, nil
}

func (s *Service) FecthCryptoDataFromAPI(cryptoName string) (models.CryptoData, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s&price_change_percentage=24h,7d,30d", cryptoName)
	resp, err := http.Get(url)
	if err != nil {
		s.Logger.Error(err.Error())
		return models.CryptoData{}, err
	}
	defer resp.Body.Close()

	var data []struct {
		ID                       string  `json:"id"`
		Symbol                   string  `json:"symbol"`
		Name                     string  `json:"name"`
		CurrentPrice             float64 `json:"current_price"`
		PriceChangePercentage24h float64 `json:"price_change_percentage_24h"`
		PriceChangePercentage7d  float64 `json:"price_change_percentage_7d_in_currency"`
		PriceChangePercentage30d float64 `json:"price_change_percentage_30d_in_currency"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		s.Logger.Error(err.Error())
		return models.CryptoData{}, err
	}

	if len(data) == 0 {
		return models.CryptoData{}, fmt.Errorf("no data found for: %s", cryptoName)
	}

	cryptoData := models.CryptoData{
		Name:                     data[0].ID,
		CurrentPrice:             data[0].CurrentPrice,
		PriceChangePercentage24h: data[0].PriceChangePercentage24h,
		PriceChangePercentage7d:  data[0].PriceChangePercentage7d,
		PriceChangePercentage30d: data[0].PriceChangePercentage30d,
		UpdatedAt:                time.Now(),
	}

	return cryptoData, nil
}
