package repository

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"wager-be/internal/domain"
)

type purchaseRepo struct {
	*contextTx
	db *gorm.DB
}

func NewPurchaseRepository(sqlDB *sql.DB) (domain.PurchaseRepository, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &purchaseRepo{
		&contextTx{
			db: gormDB,
		},
		gormDB,
	}, nil
}

func (r *purchaseRepo) Insert(ctx context.Context, p *domain.Purchase) error {
	tx := r.getTxFromContext(ctx)
	p.BoughtAt = time.Now()
	return tx.Create(p).Error
}
