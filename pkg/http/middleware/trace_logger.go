package middleware

import (
	"context"
	"products-observability/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceLoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.NewString()
		ctx := context.WithValue(c.Request.Context(), logger.TraceID, traceID)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
