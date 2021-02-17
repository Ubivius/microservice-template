package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/gorilla/mux"
)

func TestValidationMiddleware(t *testing.T) {
	productHandler := NewProductsHandler(NewTestLogger())
	// Create a router for middleware because function attachment is handled by gorilla/mux
	router := mux.NewRouter()
	router.HandleFunc("/products").Use(productHandler.MiddlewareProductValidation)

	// Creating request body
	body := &data.Product{
		Name:        "addName",
		Description: "addDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}
	bodyBytes, _ := json.Marshal(body)

	request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(bodyBytes)))
	response := httptest.NewRecorder()

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}
