package repository

import (
	"context"
	"products-observability/internal/modules/orders/model"
	"products-observability/pkg/logger"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (r *Repository) DecreaseProductStock(ctx context.Context, req model.CreateOrderRequest) error {
	var err error

	currentTime := time.Now()

	tx, ok := ctx.Value("db-tx").(*sqlx.Tx)
	if !ok {
		_, err = tx.ExecContext(ctx, decreaseProductStockQuery, req.Total, currentTime, req.ProductID)
	} else {
		_, err = r.DB.ExecContext(ctx, decreaseProductStockQuery, req.Total, currentTime, req.ProductID)
	}

	if err != nil {
		logger.Error(ctx, "failed exec context", zap.Error(err))
		return err
	}

	return nil
}
