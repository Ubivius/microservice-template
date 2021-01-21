package main

import (
	"handlers"
	"log"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "microservice-prototype", log.LstdFlags)
	hh := handlers.NewHello(l)
	gorillaMux := mux.NewRouter()
	gorillaMux.HandleFunc("/", hh.ServeHTTP)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// All that is required to run a web service
	http.ListenAndServe(":9090", sm)
}
