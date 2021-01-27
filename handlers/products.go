package handlers

import (
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

// If we don't handle specific methods, the same function is called for get, post, delete, put
func (p *Products) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// Handling get
	if request.Method == http.MethodGet {
		p.getProducts(w, request)
		return
	}
	// Handling updates

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, request *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
