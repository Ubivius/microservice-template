package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/gorilla/mux"
)

func TestValidationMiddleware(t *testing.T) {
	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()

	// Creating request body
	body := &data.Product{
		Name:        "addName",
		Description: "addDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}

	request := httptest.NewRequest(http.MethodPost, "/products", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyProduct{}, body)
	newRequest := request.WithContext(ctx)

	productHandler := NewProductsHandler(NewTestLogger())
	productHandler.AddProduct(response, newRequest)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}
