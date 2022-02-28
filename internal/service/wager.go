package service

import (
	"context"
	"math"

	"gorm.io/gorm"

	"wager-be/internal/domain"
	"wager-be/pkg/apperror"
	"wager-be/pkg/null"
)

type WagerService interface {
	CreateWager(ctx context.Context, wager *domain.Wager) error

	ListWager(ctx context.Context, pagination *domain.Pagination) ([]domain.Wager, error)
	Buy(ctx context.Context, p *domain.Purchase) error
}

type wagerService struct {
	wagerRepo    domain.WagerRepository
	purchaseRepo domain.PurchaseRepository
}

func NewWagerService(
	wagerRepo domain.WagerRepository,
	purchaseRepo domain.PurchaseRepository,
) WagerService {
	return &wagerService{
		wagerRepo:    wagerRepo,
		purchaseRepo: purchaseRepo,
	}
}

func (w *wagerService) CreateWager(ctx context.Context, wager *domain.Wager) error {
	return w.wagerRepo.Insert(ctx, wager)
}

func (w *wagerService) ListWager(ctx context.Context, pagination *domain.Pagination) ([]domain.Wager, error) {
	return w.wagerRepo.Get(ctx, pagination)
}

func (w *wagerService) Buy(ctx context.Context, p *domain.Purchase) error {
	ctxWithTx := w.wagerRepo.NewTx(ctx)

	wager, err := w.wagerRepo.GetWagerByID(ctxWithTx, p.WagerID)
	if err != nil {
		w.wagerRepo.Rollback(ctxWithTx)
		if err == gorm.ErrRecordNotFound {
			return apperror.ErrWagerNotFound
		}
		return err
	}

	if wager.CurrentSellingPrice < p.BuyingPrice {
		w.wagerRepo.Rollback(ctxWithTx)
		return apperror.ErrInvalidBuyingPrice
	}

	// update wager fields
	wager.CurrentSellingPrice = p.BuyingPrice
	if !wager.AmountSold.Valid {
		wager.AmountSold = null.Float64{
			Valid:   true,
			Float64: 0,
		}
	}

	wager.AmountSold.Float64 += float64(wager.TotalWagerValue) - wager.AmountSold.Float64
	wager.PercentageSold = null.Uint{
		Uint:  uint(math.Round(wager.AmountSold.Float64 / float64(wager.TotalWagerValue) * 100)),
		Valid: true,
	}

	if err := w.purchaseRepo.Insert(ctxWithTx, p); err != nil {
		w.wagerRepo.Rollback(ctxWithTx)
		return err
	}

	if err := w.wagerRepo.Update(ctxWithTx, wager); err != nil {
		w.wagerRepo.Rollback(ctxWithTx)
		return err
	}

	w.wagerRepo.Commit(ctxWithTx)
	return nil
}
