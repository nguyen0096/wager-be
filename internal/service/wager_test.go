package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"wager-be/internal/domain"
	_repoMock "wager-be/mocks/domain"
	"wager-be/pkg/apperror"
	"wager-be/pkg/null"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestNewWagerService(t *testing.T) {
	t.Parallel()

	wagerRepo := &_repoMock.WagerRepository{}
	purchaseRepo := &_repoMock.PurchaseRepository{}

	NewWagerService(wagerRepo, purchaseRepo)

	type args struct {
		wagerRepo    domain.WagerRepository
		purchaseRepo domain.PurchaseRepository
	}
	tests := []struct {
		name string
		args args
		want WagerService
	}{
		{
			name: "should_init_new_wager_service",
			args: args{
				wagerRepo:    wagerRepo,
				purchaseRepo: purchaseRepo,
			},
			want: &wagerService{wagerRepo, purchaseRepo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWagerService(tt.args.wagerRepo, tt.args.purchaseRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expect %v but got %v", tt.want, got)
			}
		})
	}
}

func processWager(args mock.Arguments, pos int, fn func(w *domain.Wager)) {
	if pos >= len(args) {
		return
	}
	var w *domain.Wager
	var ok bool
	if w, ok = args[pos].(*domain.Wager); ok {
		fn(w)
	}
}

func TestCreateWager(t *testing.T) {
	timeNow := time.Now()
	generatedWagerID := 1

	// functional wager repo
	wagerRepo := &_repoMock.WagerRepository{}
	wagerRepo.On("Insert", mock.Anything, mock.Anything).Run(
		func(args mock.Arguments) {
			processWager(args, 1, func(w *domain.Wager) {
				w.ID = generatedWagerID
				w.PlacedAt = timeNow
			})
		},
	).Return(nil)

	purchaseRepo := &_repoMock.PurchaseRepository{}

	type serviceArgs struct {
		wagerRepo    domain.WagerRepository
		purchaseRepo domain.PurchaseRepository
	}

	tests := []struct {
		name        string
		serviceArgs serviceArgs
		input       *domain.Wager
		want        error
	}{
		{
			name: "should_create_wager",
			serviceArgs: serviceArgs{
				wagerRepo:    wagerRepo,
				purchaseRepo: purchaseRepo,
			},
			input: &domain.Wager{
				TotalWagerValue:   100,
				Odds:              150,
				SellingPrice:      2.5,
				SellingPercentage: 100,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &wagerService{
				tt.serviceArgs.wagerRepo,
				tt.serviceArgs.purchaseRepo,
			}

			err := svc.CreateWager(context.TODO(), tt.input)
			assert.Equal(t, tt.want, err)
			assert.Equal(t, generatedWagerID, tt.input.ID)
			assert.Equal(t, timeNow, tt.input.PlacedAt)
		})
	}
}

func processPurchase(args mock.Arguments, pos int, fn func(w *domain.Purchase)) {
	if pos >= len(args) {
		return
	}
	var w *domain.Purchase
	var ok bool
	if w, ok = args[pos].(*domain.Purchase); ok {
		fn(w)
	}
}

type txCounter struct {
	commitCount   int
	rollbackCount int
}

func (t *txCounter) attachMockMethods(p *_repoMock.WagerRepository) {
	p.On("NewTx", mock.Anything).Return(context.TODO())
	p.On("Commit", mock.Anything).Run(func(args mock.Arguments) {
		t.commitCount += 1
	})
	p.On("Rollback", mock.Anything).Run(func(args mock.Arguments) {
		t.rollbackCount += 1
	})
}

func (t *txCounter) reset() {
	t.commitCount = 0
	t.rollbackCount = 0
}

type repoMethodInputAssessor struct {
	w *domain.Wager
	p *domain.Purchase
}

func (r *repoMethodInputAssessor) assertBuyUpdate(t *testing.T) {
	t.Logf("updated wager: %v", r.w.PercentageSold)
	assert.Equal(t, null.Uint{Valid: true, Uint: 100}, r.w.PercentageSold)
	assert.Equal(t, null.Float64{Valid: true, Float64: float64(r.w.TotalWagerValue)}, r.w.AmountSold)
}

func (r *repoMethodInputAssessor) reset() {
	r.w = nil
	r.p = nil
}

func TestBuy(t *testing.T) {
	timeNow := time.Now()

	generatedPurchaseID := 1
	existingWager := &domain.Wager{
		ID:                  1,
		TotalWagerValue:     100,
		Odds:                150,
		SellingPrice:        2.5,
		SellingPercentage:   50,
		CurrentSellingPrice: 2.5,
	}

	tc := &txCounter{}
	rma := &repoMethodInputAssessor{}

	// wager repo
	wagerRepo := &_repoMock.WagerRepository{}
	tc.attachMockMethods(wagerRepo)
	wagerRepo.On("GetWagerByID", mock.Anything, mock.Anything).Return(existingWager, nil)
	wagerRepo.On("Update", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		processWager(args, 1, func(w *domain.Wager) {
			rma.w = w
		})
	}).Return(nil)

	// wager repo - id not found
	wagerRepoRecordNotFound := &_repoMock.WagerRepository{}
	tc.attachMockMethods(wagerRepoRecordNotFound)
	wagerRepoRecordNotFound.On("GetWagerByID", mock.Anything, mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	// purchase repo
	purchaseRepo := &_repoMock.PurchaseRepository{}
	purchaseRepo.On("Insert", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		processPurchase(args, 1, func(p *domain.Purchase) {
			p.BoughtAt = timeNow
			p.ID = generatedPurchaseID
		})
	}).Return(nil)

	type serviceArgs struct {
		wagerRepo    domain.WagerRepository
		purchaseRepo domain.PurchaseRepository
	}

	tests := []struct {
		name        string
		serviceArgs serviceArgs
		input       *domain.Purchase
		want        error
		rollback    bool
		commit      bool
	}{
		{
			name: "given_non_existing_wager_id_should_return_err",
			serviceArgs: serviceArgs{
				wagerRepo:    wagerRepoRecordNotFound,
				purchaseRepo: purchaseRepo,
			},
			input: &domain.Purchase{
				WagerID:     existingWager.ID,
				BuyingPrice: existingWager.CurrentSellingPrice - 0.1,
			},
			want:     apperror.ErrWagerNotFound,
			rollback: true,
		},
		{
			name: "given_buying_price_bigger_than_current_selling_price_should_return_err",
			serviceArgs: serviceArgs{
				wagerRepo:    wagerRepo,
				purchaseRepo: purchaseRepo,
			},
			input: &domain.Purchase{
				WagerID:     existingWager.ID,
				BuyingPrice: existingWager.CurrentSellingPrice + 0.1,
			},
			want:     apperror.ErrInvalidBuyingPrice,
			rollback: true,
		},
		{
			name: "given_valid_buying_price_should_update_wager_and_insert_purchase",
			serviceArgs: serviceArgs{
				wagerRepo:    wagerRepo,
				purchaseRepo: purchaseRepo,
			},
			input: &domain.Purchase{
				WagerID:     existingWager.ID,
				BuyingPrice: existingWager.CurrentSellingPrice - 0.1,
			},
			want:   nil,
			commit: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc.reset()
			rma.reset()

			svc := &wagerService{
				tt.serviceArgs.wagerRepo,
				tt.serviceArgs.purchaseRepo,
			}
			err := svc.Buy(context.TODO(), tt.input)
			assert.Equal(t, tt.want, err)

			if err == nil {
				// assert wager update
				rma.p = tt.input
				rma.assertBuyUpdate(t)

				// assert purchase insert
				assert.Equal(t, generatedPurchaseID, tt.input.ID)
				assert.Equal(t, timeNow, tt.input.BoughtAt)
			}

			if tt.commit {
				assert.Equal(t, 1, tc.commitCount)
				assert.Equal(t, 0, tc.rollbackCount)
			} else if tt.rollback {
				assert.Equal(t, 1, tc.rollbackCount)
				assert.Equal(t, 0, tc.commitCount)
			}
		})
	}
}
