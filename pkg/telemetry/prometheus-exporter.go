package telemetry

import (
	"net/http"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	export "go.opentelemetry.io/otel/sdk/export/metric"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

func InitMeter() {
	config := prometheus.Config{}
	p := processor.New(
		selector.NewWithHistogramDistribution(
			histogram.WithExplicitBoundaries(config.DefaultHistogramBoundaries),
		),
		export.CumulativeExportKindSelector(),
		processor.WithMemory(true),
	)
	c := controller.New(p)
	exporter, err := prometheus.New(config, c)
	if err != nil {
		log.Error(err, "Error starting metric exporter")
	}
	global.SetMeterProvider(exporter.MeterProvider())

	http.HandleFunc("/metrics/base", exporter.ServeHTTP)
	go func() {
		_ = http.ListenAndServe(":2222", nil)
	}()

	log.Info("Prometheus exporter running on port 2222")
}
