package handler

import (
	repository "products-observability/internal/modules/orders/repository"
	usecase "products-observability/internal/modules/orders/usecase"
)

func RegisterController(r Domain) {
	repo := repository.NewOrderRepository(repository.Repository{
		NewRelic: r.NewRelic,
		DB:       r.DB,
	})

	uc := usecase.NewOrderUsecase(usecase.Usecase{
		NewRelic:   r.NewRelic,
		Repository: repo,
	})

	h := NewOrderHandler(Handler{
		NewRelic: r.NewRelic,
		Usecase:  uc,
	})

	r.Router.POST("/orders", h.CreateOrder)
}
