package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
)

// Initializes an OTLP exporter, and configures the corresponding meter provider.
func initMeterProvider(ctx context.Context, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics exporter: %w", err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(
			metric.NewPeriodicReader(
				metricExporter,
				// Default is 1m. Set to 3s for demonstrative purposes.
				metric.WithInterval(3*time.Second),
			),
		),
		metric.WithResource(res),
	)

	otel.SetMeterProvider(meterProvider)

	return meterProvider.Shutdown, nil
}
