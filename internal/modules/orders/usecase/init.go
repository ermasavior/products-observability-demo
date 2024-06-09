package usecase

import (
	"context"

	"products-observability/internal/modules/orders/model"
	repository "products-observability/internal/modules/orders/repository"
	"products-observability/internal/modules/orders/utils"
)

type Usecase struct {
	Repository repository.OrderRepository
	Telemetry  utils.Telemetry
}

type OrderUsecase interface {
	CreateOrder(context.Context, model.CreateOrderRequest) (model.Order, error)
}

func NewOrderUsecase(usecase Usecase) OrderUsecase {
	return &usecase
}
