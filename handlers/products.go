package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Ubivius/microservice-template/data"
	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object inside context
type KeyProduct struct{}

// Product handler used for getting and updating products
type ProductsHandler struct {
	logger *log.Logger
}

func NewProductsHandler(logger *log.Logger) *ProductsHandler {
	return &ProductsHandler{logger}
}

func (productHandler *ProductsHandler) UpdateProducts(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(responseWriter, "Unable to convert id to int", http.StatusBadRequest)
		return
	}

	productHandler.logger.Println("Handle PUT product", id)

	product := request.Context().Value(KeyProduct{}).(data.Product)

	// Update product
	err = data.UpdateProduct(id, &product)
	if err == data.ErrorProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(responseWriter, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (productHandler *ProductsHandler) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	productHandler.logger.Println("Handle POST Product")
	product := request.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)
}

func (productHandler *ProductsHandler) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productHandler.logger.Println("Handle GET products")
	productList := data.GetProducts()
	err := productList.ToProductJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}
