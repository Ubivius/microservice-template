package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// AddProduct creates a new product from the received JSON
func (productHandler *ProductsHandler) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	productHandler.logger.Println("Handle POST Product")
	product := &data.Product{}
	err := data.FromJSON(product, request.Body)
	if err != nil {
		productHandler.logger.Println("Error getting from json : ", err)
	}

	data.AddProduct(product)
	responseWriter.WriteHeader(200)
}
