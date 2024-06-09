package model

type CreateOrderRequest struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Total     int64 `json:"total" binding:"required"`
}
