package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// AddProduct creates a new product from the received JSON
func (productHandler *ProductsHandler) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	productHandler.logger.Println("Handle POST Product")
	productStruct := &data.Product{}
	err := data.FromJSON(productStruct, request.Body)
	if err != nil {
		productHandler.logger.Println("Error getting from json : ", err)
	}
	productHandler.logger.Println("Product struct FromJSON : ", productStruct)
	// productHandler.logger.Println(request.Context().Value(KeyProduct{}).(*data.Product))
	// product := request.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(productStruct)
	responseWriter.WriteHeader(200)
}
