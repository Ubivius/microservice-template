package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ubivius/microservice-template/pkg/handlers"
	"github.com/Ubivius/microservice-template/pkg/router"
	"go.opentelemetry.io/otel/exporters/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	// Logger
	logger := log.New(os.Stdout, "Template", log.LstdFlags)

	// Initialising open telemetry
	// Creating console exporter
	exporter, err := stdout.NewExporter(
		stdout.WithPrettyPrint(),
	)
	if err != nil {
		logger.Fatal("Failed to initialize stdout export pipeline : ", err)
	}

	// Creating tracer provider
	ctx := context.Background()
	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
	tracerProvider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(batchSpanProcessor))
	defer func() { _ = tracerProvider.Shutdown(ctx) }()

	// Creating handlers
	productHandler := handlers.NewProductsHandler(logger)

	// Router setup
	r := router.New(productHandler, logger)

	// Server setup
	server := &http.Server{
		Addr:        ":9090",
		Handler:     r,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server on port ", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			logger.Println("Error starting server : ", err)
			logger.Fatal(err)
		}
	}()

	// Handle shutdown signals from operating system
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	receivedSignal := <-signalChannel

	logger.Println("Received terminate, beginning graceful shutdown", receivedSignal)

	// Server shutdown
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_ = server.Shutdown(timeoutContext)
}
