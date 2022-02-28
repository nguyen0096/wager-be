package server

import (
	"strconv"
	"wager-be/internal/domain"
	"wager-be/pkg/apperror"

	"github.com/gin-gonic/gin"
)

type QueryKey string

const (
	PageQueryKey  QueryKey = "page"
	LimitQueryKey QueryKey = "limit"
)

func getPaginationQuery(ctx *gin.Context) (*domain.Pagination, error) {
	var err error
	pagination := &domain.Pagination{}

	page, ok := ctx.GetQuery(string(PageQueryKey))
	if !ok {
		return nil, nil
	}
	pagination.Page, err = strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	limit, ok := ctx.GetQuery(string(LimitQueryKey))
	if !ok {
		return nil, nil
	}
	pagination.Limit, err = strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	if pagination.Page == 0 || pagination.Limit == 0 {
		return nil, apperror.ErrZeroPagination
	}

	return pagination, nil
}
