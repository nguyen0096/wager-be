package postgres

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"

	"wager-be/internal/domain"
	"wager-be/pkg/config"
)

type postgres struct {
	db *sqlx.DB
}

func NewPostgresDatabase(cfg config.PostgresConfig) domain.Repository {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database))
	if err != nil {
		log.Fatalln(err)
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

	return &postgres{
		db: db,
	}
}
