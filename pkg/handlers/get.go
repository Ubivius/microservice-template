package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// GetProducts returns the full list of products
func (productHandler *ProductsHandler) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	log.Info("GetProducts request")
	productList := productHandler.db.GetProducts()
	err := json.NewEncoder(responseWriter).Encode(productList)
	if err != nil {
		log.Error(err, "Error serializing product")
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetProductByID returns a single product from the database
func (productHandler *ProductsHandler) GetProductByID(responseWriter http.ResponseWriter, request *http.Request) {
	id := getProductID(request)
	log.Info("GetProductsByID request", "id", id)

	product, err := productHandler.db.GetProductByID(id)

	switch err {
	case nil:
		err = json.NewEncoder(responseWriter).Encode(product)
		if err != nil {
			log.Error(err, "Error serializing product")
		}
	case data.ErrorProductNotFound:
		log.Error(err, "Product not found")
		http.Error(responseWriter, "Product not found", http.StatusBadRequest)
		return
	default:
		log.Error(err, "Error getting product")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}
