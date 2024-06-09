package exporter

import (
	"context"
	"products-observability/pkg/logger"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
)

func NewOTLP(endpoint string) *otlpmetric.Exporter {
	ctx := context.Background()
	metricClient := otlpmetricgrpc.NewClient(
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(endpoint))

	metricExp, err := otlpmetric.New(ctx, metricClient)
	if err != nil {
		logger.Fatal(context.Background(), "failed to create metric collector exporter")
	}

	return metricExp
}
