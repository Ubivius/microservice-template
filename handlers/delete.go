package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/data"
)

// DELETE /products/{id}
// Deletes a product with specified id from the database
func (productHandler *ProductsHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getProductId(request)
	productHandler.logger.Println("Handle DELETE product", id)

	err := data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		productHandler.logger.Println("[ERROR] deleting, id does not exist")
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		productHandler.logger.Println("[ERROR] deleting product", err)
		http.Error(responseWriter, "Erro deleting poduct", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
