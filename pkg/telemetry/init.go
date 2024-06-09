package telemetry

import (
	"context"
	"products-observability/pkg/logger"
	"products-observability/pkg/telemetry/metric"
	metricExporter "products-observability/pkg/telemetry/metric/exporter"
	"products-observability/pkg/telemetry/trace"
	traceExporter "products-observability/pkg/telemetry/trace/exporter"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
)

type Telemetry struct {
	MetricCloseFn metric.CloseFunc
	TraceCloseFn  trace.CloseFunc
}

func InitTelemetryGlobal(name, endpoint string) Telemetry {
	metricExp := metricExporter.NewOTLP(endpoint)
	pusher, pusherCloseFn, err := metric.NewMeterProviderBuilder().
		SetExporter(metricExp).
		SetHistogramBoundaries([]float64{5, 10, 25, 50, 100, 200, 400, 800, 1000}).
		Build()
	if err != nil {
		logger.Fatal(context.Background(), "failed initializing meter provider")
	}
	global.SetMeterProvider(pusher)

	spanExporter := traceExporter.NewOTLP(endpoint)
	tracerProvider, tracerProviderCloseFn, err := trace.NewTraceProviderBuilder(name).
		SetExporter(spanExporter).
		Build()
	if err != nil {
		logger.Fatal(context.Background(), "failed initializing tracer provider")
	}
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tracerProvider)

	return Telemetry{
		MetricCloseFn: pusherCloseFn,
		TraceCloseFn:  tracerProviderCloseFn,
	}
}

func ShutdownTelemetryProviders(ctx context.Context, t Telemetry) {
	err := t.MetricCloseFn(ctx)
	if err != nil {
		logger.Error(ctx, "unable to close metric provider")
	}

	err = t.TraceCloseFn(ctx)
	if err != nil {
		logger.Error(ctx, "unable to close tracer provider")
	}
}
