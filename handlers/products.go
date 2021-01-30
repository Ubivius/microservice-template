package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Ubivius/microservice-template/data"
	"github.com/gorilla/mux"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) UpdateProducts(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id to int", http.StatusBadRequest)
		return
	}

	p.logger.Println("Handle PUT product", id)

	product := request.Context().Value(KeyProduct{}).(data.Product)

	// Update product
	err = data.UpdateProduct(id, &product)
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
	p.logger.Println("Handle POST Product")
	product := request.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)
}

func (p *Products) GetProducts(w http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle GET products")
	productList := data.GetProducts()
	err := productList.ToProductJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}
