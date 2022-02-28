package domain

import (
	"context"
	"time"
	"wager-be/pkg/null"
)

type WagerRepository interface {
	ContextTx
	Insert(ctx context.Context, wager *Wager) error
	Get(ctx context.Context, pagination *Pagination) ([]Wager, error)
	Update(ctx context.Context, wager *Wager) error
	GetWagerByID(ctx context.Context, wagerID int) (*Wager, error)
}

type Wager struct {
	ID                  int          `json:"id" gorm:"id"`
	TotalWagerValue     int          `json:"total_wager_value" gorm:"total_wager_value"`
	Odds                int          `json:"odds" gorm:"odds"`
	SellingPrice        float64      `json:"selling_price" gorm:"selling_price"`
	SellingPercentage   int          `json:"selling_percentage" gorm:"selling_percentage"`
	CurrentSellingPrice float64      `json:"current_selling_price" gorm:"current_selling_price"`
	PercentageSold      null.Uint    `json:"percentage_sold" gorm:"percentage_sold"`
	AmountSold          null.Float64 `json:"amount_sold" gorm:"amount_sold"`
	PlacedAt            time.Time    `json:"placed_at" gorm:"placed_at"`
}
