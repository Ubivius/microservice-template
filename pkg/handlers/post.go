package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
	"go.opentelemetry.io/otel"
)

// AddProduct creates a new product from the received JSON
func (productHandler *ProductsHandler) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	_, span := otel.Tracer("template").Start(request.Context(), "addProduct")
	defer span.End()
	log.Info("AddProduct request")
	product := request.Context().Value(KeyProduct{}).(*data.Product)

	err := productHandler.db.AddProduct(request.Context(), product)

	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
	} else {
		responseWriter.WriteHeader(http.StatusNoContent)
	}
}
