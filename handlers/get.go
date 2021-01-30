package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/data"
)

// GET /products
// Returns the full list of products
func (productHandler *ProductsHandler) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productHandler.logger.Println("Handle GET products")
	productList := data.GetProducts()
	err := data.ToJSON(productList, responseWriter)
	if err != nil {
		productHandler.logger.Println("[ERROR] serializing product", err)
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GET /products/{id}
// Returns a single product from the database
func (productHandler *ProductsHandler) GetProductById(responseWriter http.ResponseWriter, request *http.Request) {
	id := getProductId(request)

	productHandler.logger.Println("[DEBUG] getting id", id)

	product, err := data.GetProductById(id)
	switch err {
	case nil:
	case data.ErrorProductNotFound:
		productHandler.logger.Println("[ERROR] fetching product", err)
		http.Error(responseWriter, "Product not found", http.StatusBadRequest)
		return
	default:
		productHandler.logger.Println("[ERROR] fetching product", err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(product, responseWriter)
	if err != nil {
		productHandler.logger.Println("[ERROR] serializing product", err)
	}
}
