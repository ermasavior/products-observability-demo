package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	errorPkg "products-observability/pkg/error"
	httpUtils "products-observability/pkg/http/utils"
	"products-observability/pkg/logger"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0]
		if clientError, ok := err.Err.(*errorPkg.ClientError); ok {
			logger.Error(c.Request.Context(), clientError.Raw.Error(), zap.Int("Code", clientError.Code))
			c.JSON(
				clientError.Code,
				httpUtils.NewFailedResponse(clientError.Code, clientError.Message),
			)
			return
		}

		if err.IsType(gin.ErrorTypeBind) {
			logger.Error(c.Request.Context(), err.Error(), zap.Any("Error", err))
			c.JSON(
				http.StatusBadRequest,
				httpUtils.NewFailedResponse(http.StatusBadRequest, err.Err.Error()),
			)
			return
		}

		if err.IsType(gin.ErrorTypePrivate) {
			logger.Error(c.Request.Context(), err.Error(), zap.Any("Error", err))
			c.JSON(
				http.StatusInternalServerError,
				httpUtils.NewFailedResponse(http.StatusBadRequest, httpUtils.MessageInternalServerError),
			)
			return
		}

		logger.Error(c.Request.Context(), err.Error(), zap.Any("Error", err))
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(http.StatusBadRequest, httpUtils.MessageInternalServerError),
		)
	}
}
