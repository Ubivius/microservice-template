package telemetry

import (
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func TracerProvider(url string) (*tracesdk.TracerProvider, error) {
	log.Info("Starting trace exporter")
	return nil, nil
}
