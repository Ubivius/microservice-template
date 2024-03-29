package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/data"
)

// Errors should be templated in the future.
// A good starting reference can be found here : https://github.com/nicholasjackson/building-microservices-youtube/blob/episode_7/product-api/handlers/middleware.go
// We want our validation errors to have a standard format

// MiddlewareProductValidation is used to validate incoming product JSONS
func (productHandler *ProductsHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		product := &data.Product{}

		err := json.NewDecoder(request.Body).Decode(product)
		if err != nil {
			log.Error(err, "Error deserializing product")
			http.Error(responseWriter, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = product.ValidateProduct()
		if err != nil {
			log.Error(err, "Error validating product")
			http.Error(responseWriter, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// Add the product to the context
		ctx := context.WithValue(request.Context(), KeyProduct{}, product)
		request = request.WithContext(ctx)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(responseWriter, request)
	})
}
