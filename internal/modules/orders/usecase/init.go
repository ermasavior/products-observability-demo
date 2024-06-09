package usecase

import (
	"context"

	"products-observability/internal/modules/orders/model"
	repository "products-observability/internal/modules/orders/repository"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type Usecase struct {
	NewRelic   *newrelic.Application
	Repository repository.OrderRepository
}

type OrderUsecase interface {
	CreateOrder(context.Context, model.CreateOrderRequest) (model.Order, error)
}

func NewOrderUsecase(usecase Usecase) OrderUsecase {
	return &usecase
}
