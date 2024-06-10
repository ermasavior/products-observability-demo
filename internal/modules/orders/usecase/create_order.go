package usecase

import (
	"context"
	"database/sql"

	"products-observability/internal/modules/orders/model"
	"products-observability/pkg/logger"

	"go.uber.org/zap"
)

func (u *Usecase) CreateOrder(ctx context.Context, req model.CreateOrderRequest) (model.Order, error) {
	ctx, span := u.Telemetry.Tracer.Start(ctx, "orders.CreateOrder")
	defer span.End()

	var err error

	ctx, err = u.Repository.BeginTx(ctx)
	if err != nil {
		logger.Error(ctx, "error begin tx", zap.Int64("product_id", req.ProductID))
		return model.Order{}, err
	}
	defer u.Repository.RollbackTx(ctx)

	stock, err := u.Repository.GetProductStockByID(ctx, req.ProductID)
	if err == sql.ErrNoRows {
		return model.Order{}, model.ErrorProductNotFound
	}
	if err != nil {
		logger.Error(ctx, "error get product stock by id", zap.Int64("product_id", req.ProductID))
		return model.Order{}, err
	}

	if req.Total > stock {
		return model.Order{}, model.ErrorProductNotEnoughStock
	}

	err = u.Repository.DecreaseProductStock(ctx, req)
	if err != nil {
		logger.Error(ctx, "error decrease product stock", zap.Int64("product_id", req.ProductID))
		return model.Order{}, err
	}

	res, err := u.Repository.InsertNewOrder(ctx, req)
	if err != nil {
		logger.Error(ctx, "error insert new order", zap.Int64("product_id", req.ProductID))
		return model.Order{}, err
	}

	err = u.Repository.FinishTx(ctx)
	if err != nil {
		logger.Error(ctx, "error finish tx", zap.Int64("product_id", req.ProductID))
		return model.Order{}, err
	}

	u.Telemetry.OrderCounter.Add(ctx, 1)

	return res, nil
}
