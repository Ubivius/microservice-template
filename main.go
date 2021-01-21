package main

import (
	"log"
	"net/http"
)

func main() {
	// Handle func is a convenience method on the http package.
	// Registers a function to a path on the default serve mux (http handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye")
	})

	// All that is required to run a web service
	http.ListenAndServe(":9090", nil)
}
