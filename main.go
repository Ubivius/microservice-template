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
	// hh := handlers.NewHello(l)
	achievementHandlers := handlers.NewAchievement(l)
	gorillaMux := mux.NewRouter()
	gorillaMux.HandleFunc("/", handlers.NewHello(l).ServeHTTP)
	gorillaMux.HandleFunc("/achievement", achievementHandlers.ServeHTTP)

	// All that is required to run a web service
	http.ListenAndServe(":9090", gorillaMux)
}
