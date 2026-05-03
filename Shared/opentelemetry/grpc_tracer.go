package opentelemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

func NewGrpcTracer(ctx context.Context, serviceName, host, port string) (*sdktrace.TracerProvider, error) {

	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(host+":"+port),
	)
	if err != nil {
		return nil, err
	}
	sampler := sdktrace.TraceIDRatioBased(0.3) //30% of spans are processed

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter), //Processing spans in batches
		sdktrace.WithSampler(sampler),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceName(serviceName),
			),
		),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}
