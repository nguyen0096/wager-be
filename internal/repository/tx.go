package repository

import (
	"context"

	"gorm.io/gorm"
)

type contextTxKey string

const (
	TxContextTxKey contextTxKey = "db_context"
)

type contextTx struct {
	db *gorm.DB
}

func (dc *contextTx) NewTx(ctx context.Context) context.Context {
	tx := dc.db.Begin()
	ctxWithTx := context.WithValue(ctx, TxContextTxKey, tx)
	return ctxWithTx
}

func (dc *contextTx) getTxFromContext(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(TxContextTxKey).(*gorm.DB)
	if !ok {
		return dc.db
	}
	return tx
}

func (dc *contextTx) Commit(ctx context.Context) {
	tx := dc.getTxFromContext(ctx)
	tx.Commit()
}

func (dc *contextTx) Rollback(ctx context.Context) {
	tx := dc.getTxFromContext(ctx)
	tx.Rollback()
}
