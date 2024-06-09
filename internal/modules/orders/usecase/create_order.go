package usecase

import (
	"context"
	"database/sql"

	"products-observability/internal/modules/orders/model"
)

func (u *Usecase) CreateOrder(ctx context.Context, req model.CreateOrderRequest) (model.Order, error) {
	var err error

	ctx, err = u.Repository.BeginTx(ctx)
	if err != nil {
		return model.Order{}, err
	}
	defer u.Repository.RollbackTx(ctx)

	stock, err := u.Repository.GetProductStockByID(ctx, req.ProductID)
	if err == sql.ErrNoRows {
		return model.Order{}, model.ErrorProductNotFound
	}
	if err != nil {
		return model.Order{}, err
	}

	if req.Total > stock {
		return model.Order{}, model.ErrorProductNotEnoughStock
	}

	res, err := u.Repository.InsertNewOrder(ctx, req)
	if err != nil {
		return model.Order{}, err
	}

	err = u.Repository.DecreaseProductStock(ctx, req)
	if err != nil {
		return model.Order{}, err
	}

	u.Repository.FinishTx(ctx)

	return res, nil
}
