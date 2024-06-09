package repository

import (
	"context"
	"time"

	"products-observability/internal/modules/orders/model"
	"products-observability/pkg/logger"
	"products-observability/pkg/utils"

	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

func (r *Repository) InsertNewOrder(ctx context.Context, req model.CreateOrderRequest) (model.Order, error) {
	txn := r.NewRelic.StartTransaction("InsertNewOrder")
	defer txn.End()
	ctx = newrelic.NewContext(ctx, txn)

	var (
		err  error
		rows *sqlx.Rows
	)

	currentTime := time.Now()

	newOrder := model.Order{
		ProductID: req.ProductID,
		Total:     req.Total,
		CreatedAt: currentTime.Format(time.RFC3339),
	}

	tx, ok := ctx.Value(utils.DBTxType).(*sqlx.Tx)
	if !ok {
		rows, err = r.DB.NamedQuery(insertOrderQuery, newOrder)
	} else {
		rows, err = tx.NamedQuery(insertOrderQuery, newOrder)
	}

	if err != nil {
		logger.Error(ctx, "failed exec context", zap.Error(err))
		return model.Order{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&newOrder); err != nil {
			logger.Error(ctx, "failed struct scan", zap.Error(err))
		}
	}

	return newOrder, nil
}
