package repository

import (
	"context"
	"database/sql"
	"errors"
	"products-observability/pkg/logger"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (r *Repository) BeginTx(ctx context.Context) (context.Context, error) {
	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, "db-tx", tx)
	return ctx, nil
}

func (r *Repository) FinishTx(ctx context.Context) error {
	tx, ok := ctx.Value("db-tx").(*sqlx.Tx)
	if !ok {
		return errors.New("tx not found")
	}

	if err := tx.Commit(); err != nil {
		logger.Error(ctx, "error commit tx", zap.Error(err))
		r.RollbackTx(ctx)
		return err
	}
	return nil
}

func (r *Repository) RollbackTx(ctx context.Context) error {
	tx, ok := ctx.Value("db-tx").(*sqlx.Tx)
	if !ok {
		return errors.New("tx not found")
	}

	err := tx.Rollback()
	if err == sql.ErrTxDone {
		return nil
	}
	if err != nil {
		logger.Error(ctx, "error rollback", zap.Error(err))
		return err
	}

	return nil
}
