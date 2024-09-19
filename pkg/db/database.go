package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/skinkvi/crpt/internal/config"
	"github.com/skinkvi/crpt/pkg/util/logger"
)

func Init() *sqlx.DB {
	dbConfig := config.GetConfig().DB

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname, dbConfig.Sslmode)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}

	if err := db.Ping(); err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}

	return db
}
