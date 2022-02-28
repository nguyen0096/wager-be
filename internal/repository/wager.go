package repository

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"wager-be/internal/domain"
)

type wagerRepo struct {
	*contextTx
	db *gorm.DB
}

func NewWagerRepository(sqlDB *sql.DB) (domain.WagerRepository, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &wagerRepo{
		&contextTx{
			db: gormDB,
		},
		gormDB,
	}, nil
}

func (r *wagerRepo) Insert(ctx context.Context, w *domain.Wager) error {
	tx := r.getTxFromContext(ctx)
	w.PlacedAt = time.Now()
	w.CurrentSellingPrice = w.SellingPrice
	return tx.Create(w).Error
}

func (r *wagerRepo) Update(ctx context.Context, w *domain.Wager) error {
	tx := r.getTxFromContext(ctx)
	return tx.Save(w).Error
}

func (r *wagerRepo) Get(ctx context.Context, pagination *domain.Pagination) ([]domain.Wager, error) {
	tx := r.getTxFromContext(ctx)
	wagers := []domain.Wager{}

	tx = tx.Model(&domain.Wager{})
	if pagination != nil {
		tx.Limit(pagination.Limit)
		tx.Offset((pagination.Page - 1) * pagination.Limit)
	}

	err := tx.Order("placed_at desc").Find(&wagers).Error
	return wagers, err
}

func (r *wagerRepo) GetWagerByID(ctx context.Context, wagerID int) (*domain.Wager, error) {
	wager := domain.Wager{}
	tx := r.getTxFromContext(ctx)
	err := tx.Model(&domain.Wager{}).Where("id = ?", wagerID).First(&wager).Error
	return &wager, err
}
