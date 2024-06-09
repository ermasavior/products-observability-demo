package handler

import (
	"net/http"
	"products-observability/internal/modules/orders/model"
	httpUtils "products-observability/pkg/http/utils"
	"products-observability/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	var req model.CreateOrderRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, utils.ValidateRequestStruct(err).Error()),
		)
		return
	}

	data, errUC := h.Usecase.CreateOrder(c.Request.Context(), req)
	if errUC == model.ErrorProductNotEnoughStock || errUC == model.ErrorProductNotFound {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, errUC.Error()),
		)
		return
	}

	if errUC != nil {
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(http.StatusInternalServerError, errUC.Error()),
		)
		return
	}

	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(data))
}
