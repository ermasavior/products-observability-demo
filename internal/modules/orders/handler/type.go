package handler

import (
	usecase "products-observability/internal/modules/orders/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Domain struct {
	Router   *gin.Engine
	NewRelic *newrelic.Application
	DB       *sqlx.DB
}

type Handler struct {
	NewRelic *newrelic.Application
	Usecase  usecase.OrderUsecase
}

type OrderHandler interface {
	CreateOrder(c *gin.Context)
}

func NewOrderHandler(h Handler) OrderHandler {
	return &h
}
