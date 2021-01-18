package main

import (
	"log"
	"net/http"
)

func main() {
	// Handle func is a convenience method on the http package
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})
	// All that is required to run a web service
	http.ListenAndServe(":9090", nil)
}
