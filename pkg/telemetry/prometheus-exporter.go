package telemetry

import (
	"net/http"

	"contrib.go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

var TotalConnections = stats.Int64("connections", "The number of connections to the service", "{tot}")

var connectionCountView = &view.View{
	Name:        "total_connection_count",
	Measure:     TotalConnections,
	Description: "The number of connections to the service",
	Aggregation: view.Count(),
}

func StartPrometheusExporterWithName(exporterNamespace string) {
	// Register the view. It is imperative that this step exists,
	// otherwise recorded metrics will be dropped and never exported.
	if err := view.Register(connectionCountView); err != nil {
		log.Error(err, "Failed to register OpenCensus metric views")
	}

	// Create the prometheus exporter
	prometheusExporter, err := prometheus.NewExporter(prometheus.Options{
		Namespace: exporterNamespace,
	})
	if err != nil {
		log.Error(err, "Failed to create the Prometheus stats exporter")
	}

	// Now finally run the Prometheus exporter as a scrape endpoint.
	// We'll run the server on port 8888.
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", prometheusExporter)
		if err := http.ListenAndServe(":8888", mux); err != nil {
			log.Error(err, "Failed to run Prometheus scrape endpoint")
		}
	}()
}
