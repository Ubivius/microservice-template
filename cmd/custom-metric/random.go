package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

var latencyMs = stats.Float64("task_latency", "The task latency in milliseconds", "ms")
var totalConnections = stats.Int64("connections", "The number of connections to the service", "{tot}")

func main() {
	// Register the view. It is imperative that this step exists,
	// otherwise recorded metrics will be dropped and never exported.
	latencyView := &view.View{
		Name:        "task_latency_distribution",
		Measure:     latencyMs,
		Description: "The distribution of the task latencies",

		// Latency in buckets:
		// [>=0ms, >=100ms, >=200ms, >=400ms, >=1s, >=2s, >=4s]
		Aggregation: view.Distribution(0, 100, 200, 400, 1000, 2000, 4000),
	}

	connectionCountView := &view.View{
		Name:        "total_connection_count",
		Measure:     totalConnections,
		Description: "The number of connections to the service",
		Aggregation: view.Count(),
	}

	if err := view.Register(latencyView, connectionCountView); err != nil {
		log.Fatalf("Failed to register the view: %v", err)
	}

	// Create the Prometheus exporter.
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "golangsvc",
	})
	if err != nil {
		log.Fatalf("Failed to create the Prometheus stats exporter: %v", err)
	}

	// Now finally run the Prometheus exporter as a scrape endpoint.
	// We'll run the server on port 8888.
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", pe)
		if err := http.ListenAndServe(":8888", mux); err != nil {
			log.Fatalf("Failed to run Prometheus scrape endpoint: %v", err)
		}
	}()

	for i := 0; i < 100; i++ {
		ms := float64(5*time.Second/time.Millisecond) * rand.Float64()
		fmt.Printf("Latency %d: %f\n", i, ms)
		stats.Record(context.Background(), latencyMs.M(ms), totalConnections.M(1))
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Done recording metrics")
}
