package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// UpdateProducts updates the product with the ID specified in the received JSON product
func (productHandler *ProductsHandler) UpdateProducts(responseWriter http.ResponseWriter, request *http.Request) {
	product := request.Context().Value(KeyProduct{}).(*data.Product)
	log.Info("UpdateProducts request", "id", product.ID)

	// Update product
	err := productHandler.db.UpdateProduct(product)
	if err == data.ErrorProductNotFound {
		log.Error(err, "Product not found")
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	// Returns status, no content required
	responseWriter.WriteHeader(http.StatusNoContent)
}
