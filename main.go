package main

import (
	"bytes"
	"handlers"
	"log"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch"
	"github.com/gorilla/mux"
)

func main() {
	//Logger
	l := log.New(os.Stdout, "microservice-prototype", log.LstdFlags)

	// Configuration elastic search
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		l.Println("Error creating the es client.")
	}

	var b bytes.Buffer
	b.WriteString(`{"Users" : "Jeremi"}`)

	res, _ := es.Index("method1", &b)
	l.Println(res)

	// Handlers
	helloHandler := handlers.NewHello(l)
	achievementHandlers := handlers.NewAchievement(l)

	// Routing
	gorillaMux := mux.NewRouter()
	gorillaMux.HandleFunc("/", helloHandler.ServeHTTP)
	gorillaMux.HandleFunc("/achievement", achievementHandlers.ServeHTTP)

	// Start server
	http.ListenAndServe(":9090", gorillaMux)
}
