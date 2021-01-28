package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Ubivius/microservice-template/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) UpdateProducts(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id to int", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT product", id)
	product := &data.Product{}

	// Parsing to json
	err = product.FromJson(request.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	// Update product
	err = data.UpdateProduct(id, product)
	if err == data.ErrorProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle POST product")
	product := &data.Product{}
	err := product.FromJson(request.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(product)
}

func (p *Products) GetProducts(w http.ResponseWriter, request *http.Request) {
	p.l.Println("Handle GET products")
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
