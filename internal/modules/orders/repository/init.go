package repository

import (
	"context"

	"products-observability/internal/modules/orders/model"

	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Repository struct {
	NewRelic *newrelic.Application
	DB       *sqlx.DB
}

type OrderRepository interface {
	BeginTx(ctx context.Context) (context.Context, error)
	FinishTx(ctx context.Context) error
	RollbackTx(ctx context.Context) error

	GetProductStockByID(ctx context.Context, productID int64) (int64, error)
	InsertNewOrder(context.Context, model.CreateOrderRequest) (model.Order, error)
	DecreaseProductStock(ctx context.Context, req model.CreateOrderRequest) error
}

func NewOrderRepository(repo Repository) OrderRepository {
	return &repo
}
