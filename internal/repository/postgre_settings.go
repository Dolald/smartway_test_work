package repository

import (
	"fmt"

	"github.com/Dolald/smartway_test_work/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *configs.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("postgre opening failed")
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db ping failed")
	}

	return db, nil
}
