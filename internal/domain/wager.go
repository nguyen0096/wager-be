package domain

import "time"

type Wager struct {
	ID                  int
	TotalWagerValue     int
	Odds                int
	SellingPrice        int
	SellingPercentage   float64
	CurrentSellingPrice float64
	PercentageSold      float64
	AmountSold          int
	PlacedAt            time.Time
}
