package models

import "time"

type CryptoData struct {
	Name                     string    `db:"name" json:"name"`
	CurrentPrice             float64   `db:"current_price" json:"current_price"`
	PriceChangePercentage24h float64   `db:"price_change_24h" json:"price_change_24h"`
	PriceChangePercentage7d  float64   `db:"price_change_7d" json:"price_change_7d"`
	PriceChangePercentage30d float64   `db:"price_change_30d" json:"price_change_30d"`
	UpdatedAt                time.Time `db:"updated_at" json:"updated_at"`
}
