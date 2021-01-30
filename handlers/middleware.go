package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ubivius/microservice-template/data"
)

// Json Product Validation
func (productHandler *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		product := &data.Product{}

		err := data.FromJSON(product, request.Body)
		if err != nil {
			productHandler.logger.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = product.ValidateProduct()
		if err != nil {
			productHandler.logger.Println("[ERROR] validating product", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// Add the product to the context
		context := context.WithValue(request.Context(), KeyProduct{}, product)
		newRequest := request.WithContext(context)

		// Call the next handler, which can be another middleware or the final handler
		next.ServeHTTP(w, newRequest)
	})
}
