package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Handle func is a convenience method on the http package.
	// Registers a function to a path on the default serve mux (http handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data : %s\n", data)
		fmt.Fprintf(w, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye")
	})

	// All that is required to run a web service
	http.ListenAndServe(":9090", nil)
}
