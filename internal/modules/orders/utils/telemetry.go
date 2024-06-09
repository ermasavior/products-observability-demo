package utils

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/unit"
	"go.opentelemetry.io/otel/trace"
)

type Telemetry struct {
	Tracer       trace.Tracer
	Meter        metric.Meter
	OrderCounter syncint64.Counter
}

func InitTelemetry() Telemetry {
	meter := global.Meter("products-observability/orders")
	orderCounter, _ := meter.SyncInt64().Counter(
		"order",
		instrument.WithDescription("number of orders"),
		instrument.WithUnit(unit.Dimensionless))

	return Telemetry{
		Tracer:       otel.Tracer("products-observability/orders"),
		Meter:        meter,
		OrderCounter: orderCounter,
	}
}
