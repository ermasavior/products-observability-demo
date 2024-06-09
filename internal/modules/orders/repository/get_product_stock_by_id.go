package repository

import (
	"context"
	"products-observability/pkg/logger"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (r *Repository) GetProductStockByID(ctx context.Context, productID int64) (int64, error) {
	var (
		stock int64
		err   error
	)

	tx, ok := ctx.Value("db-tx").(*sqlx.Tx)
	if !ok {
		err = r.DB.GetContext(ctx, &stock, getProductStockByIDQuery, productID)
	} else {
		err = tx.GetContext(ctx, &stock, getProductStockByIDQuery, productID)
	}

	if err != nil {
		logger.Error(ctx, "failed get context", zap.Error(err))
		return 0, err
	}

	return stock, nil
}
