package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// AddProduct creates a new product from the received JSON
func (productHandler *ProductsHandler) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	// productHandler.logger.Println("Handle POST Product")
	product := request.Context().Value(KeyProduct{}).(*data.Product)

	err := productHandler.db.AddProduct(product)

	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		responseWriter.WriteHeader(http.StatusNoContent)
	}
}
