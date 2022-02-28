package domain

import (
	"context"
	"time"
)

type PurchaseRepository interface {
	Insert(ctx context.Context, purchase *Purchase) error
}

type Purchase struct {
	ID          int       `json:"id" gorm:"id"`
	WagerID     int       `json:"wager_id" gorm:"id"`
	BuyingPrice float64   `json:"buying_price"`
	BoughtAt    time.Time `json:"bought_at"`
}
