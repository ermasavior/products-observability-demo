package http

import (
	"products-observability/pkg/http/middleware"
	"products-observability/pkg/telemetry/otelginmetrics"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// NewHTTPServer returns gin http server
func NewHTTPServer(appName, appEnv string) *gin.Engine {
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	if ve, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ve.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})
	}

	gin.EnableJsonDecoderDisallowUnknownFields()

	router := gin.New()

	router.Use(middleware.LogHandler())
	router.Use(gin.Recovery())
	router.Use(middleware.CorsHandler())
	router.Use(middleware.TraceLoggerHandler())
	router.Use(otelgin.Middleware(appName))
	router.Use(otelginmetrics.Middleware(appName))

	return router
}
