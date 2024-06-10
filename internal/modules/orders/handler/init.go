package handler

import (
	repository "products-observability/internal/modules/orders/repository"
	usecase "products-observability/internal/modules/orders/usecase"
	"products-observability/internal/modules/orders/utils"
)

func RegisterController(r Domain) {
	telemetry := utils.InitTelemetry()
	repo := repository.NewOrderRepository(repository.Repository{
		NewRelic: r.NewRelic,
		DB:       r.DB,
	})

	uc := usecase.NewOrderUsecase(usecase.Usecase{
		Repository: repo,
		Telemetry:  telemetry,
	})

	h := NewOrderHandler(Handler{
		NewRelic: r.NewRelic,
		Usecase:  uc,
	})

	// orderHandler := otelhttp.NewHandler(http.HandlerFunc(h.CreateOrder), "products")
	r.Router.POST("/orders", h.CreateOrder)
}
