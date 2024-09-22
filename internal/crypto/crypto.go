// pkg/crypto/crypto.go
package crypto

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/skinkvi/crpt/internal/models"
	"github.com/skinkvi/crpt/pkg/util/logger"
)

func GetCryptoData(cryptoName string) (models.CryptoData, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd&include_24hr_change=true&include_7d_change=true&include_30d_change=true", cryptoName)
	resp, err := http.Get(url)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return models.CryptoData{}, err
	}
	defer resp.Body.Close()

	var data map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		logger.GetLogger().Error(err.Error())
		return models.CryptoData{}, err
	}

	if data == nil {
		return models.CryptoData{}, fmt.Errorf("error parsing coingecko data: data is nil")
	}

	if data[cryptoName] == nil {
		return models.CryptoData{}, fmt.Errorf("error parsing coingecko data: missing %s data", cryptoName)
	}

	cryptoData := models.CryptoData{
		Name:           cryptoName,
		CurrentPrice:   data[cryptoName]["usd"],
		PriceChange4h:  data[cryptoName]["usd_4h_change"],
		PriceChange24h: data[cryptoName]["usd_24h_change"],
		PriceChange7d:  data[cryptoName]["usd_7d_change"],
	}

	return cryptoData, nil
}
