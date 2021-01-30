package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/data"
)

func (productHandler *ProductsHandler) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productHandler.logger.Println("Handle GET products")
	productList := data.GetProducts()
	err := productList.ToProductJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal json", http.StatusInternalServerError)
	}
}
