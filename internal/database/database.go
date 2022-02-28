package database

import (
	"database/sql"
	"fmt"
	"time"
	"wager-be/pkg/config"

	_ "github.com/lib/pq"
)

func NewPostgresConn(cfg config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database))
	if err != nil {
		return nil, err
	}
	if cfg.MaxConn > 0 {
		db.SetMaxOpenConns(cfg.MaxConn)
	}
	if cfg.MaxIdleConn > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConn)
	}
	if cfg.MaxIdleConnTime > 0 {
		db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleConnTime))
	}
	return db, nil
}
