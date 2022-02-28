package apperror

import "fmt"

var (
	ErrZeroPagination       = fmt.Errorf("pagination cannot be zero")
	ErrInvalidRequestFormat = fmt.Errorf("request body has invalid format")
	ErrInternalServer       = fmt.Errorf("unknown error")
	ErrWagerNotFound        = fmt.Errorf("wager not found")
	ErrInvalidBuyingPrice   = fmt.Errorf("buying price must be smaller or equal to current selling price of wager")
)
