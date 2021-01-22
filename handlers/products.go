package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ubivius/microservice-template/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, h *http.Request) {
	productList := data.GetProducts()
	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	w.Write(data)
}
