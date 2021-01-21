package main

import (
	"example/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// All that is required to run a web service
	http.ListenAndServe(":9090", sm)
}
