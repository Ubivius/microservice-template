package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// Delete a product with specified id from the database
func (productHandler *ProductsHandler) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getProductID(request)
	log.Info("Delete product by ID request", "id", id)

	err := productHandler.db.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		log.Error(err, "Error deleting product, id does not exist")
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Error(err, "Error deleting product")
		http.Error(responseWriter, "Error deleting poduct", http.StatusInternalServerError)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
