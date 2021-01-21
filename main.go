package main

import (
	"handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	l := log.New(os.Stdout, "microservice-prototype", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	achievementHandlers := handlers.NewAchievement(l)
	gorillaMux := mux.NewRouter()
	gorillaMux.HandleFunc("/", helloHandler.ServeHTTP)
	gorillaMux.HandleFunc("/achievement", achievementHandlers.ServeHTTP)

	// All that is required to run a web service
	http.ListenAndServe(":9090", gorillaMux)
}
