package domain

import "context"

type ContextTx interface {
	NewTx(ctx context.Context) context.Context
	Commit(ctx context.Context)
	Rollback(ctx context.Context)
}
