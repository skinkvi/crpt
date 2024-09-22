package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/skinkvi/crpt/internal/config"
	"github.com/skinkvi/crpt/pkg/util/logger"
)

func Init() *sqlx.DB {
	dbConfig := config.GetConfig().DB

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname, dbConfig.Sslmode)

	var db *sqlx.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sqlx.Open("postgres", connStr)
		if err != nil {
			logger.GetLogger().Error(err.Error())
			time.Sleep(2 * time.Second)
			continue
		}

		if err := db.Ping(); err != nil {
			logger.GetLogger().Error(err.Error())
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}

	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}

	return db
}
