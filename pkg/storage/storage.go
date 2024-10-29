package storage

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/skinkvi/crpt/internal/config"
	"github.com/skinkvi/crpt/internal/models"
	"go.uber.org/zap"
)

type Storage interface {
	GetCryptoData(name string) (models.CryptoData, error)
	SaveCryptoData(data models.CryptoData) error
}

type DBStorage struct {
	DB     *sqlx.DB
	Logger *zap.Logger
}

func NewStorage(logger *zap.Logger) *DBStorage {
	dbConfig := config.GetConfig().DB

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Dbname,
		dbConfig.Sslmode)

	var db *sqlx.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sqlx.Open("postgres", connString)
		if err != nil {
			logger.Sugar().Error("Failed to connect to database: ", err.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		if err := db.Ping(); err != nil {
			logger.Sugar().Error("Failed to Ping database: ", err.Error())
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}

	if err != nil {
		logger.Sugar().Fatal("Could not connect to database:", err)
	}

	logger.Info("Successful connection to database")

	return &DBStorage{
		DB:     db,
		Logger: logger,
	}
}

func (s *DBStorage) GetCryptoData(name string) (models.CryptoData, error) {
	var data models.CryptoData
	err := s.DB.Get(&data, "SELECT * FROM crypto_data WHERE name = $1", name)
	if err != nil {
		return models.CryptoData{}, err
	}

	return data, nil
}

func (s *DBStorage) SaveCryptoData(data models.CryptoData) error {
	_, err := s.DB.Exec(`
        INSERT INTO crypto_data (name, current_price, price_change_24h, price_change_7d, price_change_30d, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (name) DO UPDATE SET
            current_price = EXCLUDED.current_price,
            price_change_24h = EXCLUDED.price_change_24h,
            price_change_7d = EXCLUDED.price_change_7d,
            price_change_30d = EXCLUDED.price_change_30d,
            updated_at = EXCLUDED.updated_at
    `,
		data.Name,
		data.CurrentPrice,
		data.PriceChangePercentage24h,
		data.PriceChangePercentage7d,
		data.PriceChangePercentage30d,
		data.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
