package main

import (
	"handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Creating handlers
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	// Mux route handling with default http ServeMux
	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)

	// Server setup
	server := &http.Server{
		Addr:        ":9090",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
	}
	server.ListenAndServe()
}
