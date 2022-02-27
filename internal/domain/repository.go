package domain

import "context"

type Repository interface {
	InsertWager(ctx context.Context, wager *Wager) error
}
