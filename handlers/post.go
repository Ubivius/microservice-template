package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/data"
)

// /POST /products
// Creates a new product
func (productHandler *ProductsHandler) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	productHandler.logger.Println("Handle POST Product")
	product := request.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(product)
}
