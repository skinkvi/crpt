// internal/models/models.go
package models

type CryptoData struct {
	Name           string  `json:"name"`
	CurrentPrice   float64 `json:"current_price"`
	PriceChange4h  float64 `json:"price_change_4h"`
	PriceChange24h float64 `json:"price_change_24h"`
	PriceChange7d  float64 `json:"price_change_7d"`
}
