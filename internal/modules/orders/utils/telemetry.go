package utils

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Telemetry struct {
	Tracer       trace.Tracer
	Meter        metric.Meter
	OrderCounter metric.Int64Counter
}

func InitTelemetry() Telemetry {
	meter := otel.Meter("products-observability/orders")
	orderCounter, _ := meter.Int64Counter(
		"order_counter",
		metric.WithDescription("number of orders"),
		metric.WithUnit("{orders}"))

	return Telemetry{
		Tracer:       otel.Tracer("products-observability/orders"),
		Meter:        meter,
		OrderCounter: orderCounter,
	}
}
