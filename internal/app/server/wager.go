package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"wager-be/internal/domain"
	"wager-be/pkg/apperror"
	"wager-be/pkg/validator"
)

type createWagerRequest struct {
	TotalWagerValue   int     `json:"total_wager_value" validate:"required,gt=0"`
	Odds              int     `json:"odds" validate:"required"`
	SellingPercentage int     `json:"selling_percentage" validate:"required"`
	SellingPrice      float64 `json:"selling_price" validate:"required"`
}

func (s *server) handleCreateWager(ctx *gin.Context) {
	req := createWagerRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		s.logger.Error(err, "failed to bind json")
		s.responseErr(ctx, http.StatusBadRequest, apperror.ErrInvalidRequestFormat)
		return
	}

	if err := validator.Get().Struct(req); err != nil {
		s.responseErr(ctx, http.StatusBadRequest, err)
		return
	}

	wager := &domain.Wager{
		TotalWagerValue:   req.TotalWagerValue,
		Odds:              req.Odds,
		SellingPrice:      req.SellingPrice,
		SellingPercentage: req.SellingPercentage,
	}

	err = s.wagerService.CreateWager(ctx, wager)
	if err != nil {
		s.logger.Error(err, "failed to create wager")
		s.responseErr(ctx, http.StatusInternalServerError, apperror.ErrInvalidRequestFormat)
		return
	}

	ctx.JSON(http.StatusCreated, wager)
}

func (s *server) handleListWager(ctx *gin.Context) {
	pagination, err := getPaginationQuery(ctx)
	if err != nil {
		s.logger.Error(err, "failed to parse pagination query")
		s.responseErr(ctx, http.StatusBadRequest, err)
		return
	}

	wagers, err := s.wagerService.ListWager(ctx, pagination)
	if err != nil {
		s.logger.Error(err, "failed to list wager")
		s.responseErr(ctx, http.StatusInternalServerError, apperror.ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK, wagers)
}

type buyRequest struct {
	BuyingPrice float64 `json:"buying_price"`
}

func (s *server) handleBuy(ctx *gin.Context) {
	wagerIDStr, ok := ctx.Params.Get("wager_id")
	if !ok {
		s.responseErr(ctx, http.StatusBadRequest, apperror.ErrInvalidRequestFormat)
		return
	}

	wagerID, err := strconv.Atoi(wagerIDStr)
	if err != nil {
		s.responseErr(ctx, http.StatusBadRequest, apperror.ErrInvalidRequestFormat)
		return
	}

	req := buyRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		s.logger.Error(err, "failed to bind json")
		s.responseErr(ctx, http.StatusBadRequest, apperror.ErrInvalidRequestFormat)
		return
	}

	p := &domain.Purchase{
		WagerID:     wagerID,
		BuyingPrice: req.BuyingPrice,
	}

	err = s.wagerService.Buy(ctx, p)
	if err != nil {
		s.logger.Error(err, "failed to buy")
		s.responseErr(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, p)
}
