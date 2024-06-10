package telemetry

import (
	"context"
	"errors"
	"fmt"
	"products-observability/pkg/logger"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// SetupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func SetupOTelSDK(ctx context.Context, appName, otelAddress string) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// shutdown will flush any remaining spans and shut down the exporter.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			semconv.ServiceNameKey.String(appName),
		),
	)
	if err != nil {
		logger.Fatal(ctx, "error init otel resource")
	}

	conn, err := initConn(otelAddress)
	if err != nil {
		logger.Fatal(ctx, "error init otel connection")
	}

	// Set up propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Set up trace provider.
	tracerProviderFn, err := initTracerProvider(ctx, res, conn)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProviderFn)

	// Set up meter provider.
	meterProviderFn, err := initMeterProvider(ctx, res, conn)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProviderFn)

	return
}

func initConn(otelAddress string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		otelAddress,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, err
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
