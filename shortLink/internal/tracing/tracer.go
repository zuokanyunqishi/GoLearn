package tracing

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracerProvider(name string, version string, enabled bool, ratio any) (*sdktrace.TracerProvider, error) {

	if enabled == true {
	}

	return nil, nil
}
