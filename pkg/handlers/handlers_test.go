package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/Ubivius/microservice-template/pkg/database"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func NewProductDB() database.ProductDB {
	return database.NewMockProducts()
}

func TestGetProducts(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewProductDB())
	productHandler.GetProducts(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}

	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") || !strings.Contains(response.Body.String(), "e2382ea2-b5fa-4506-aa9d-d338aa52af44") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetExistingProductByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewProductDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetProductByID(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got : %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "a2181017-5c53-422b-b6bc-036b27c04fc8") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetNonExistingProductByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products/4", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewProductDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetProductByID(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got : %d", http.StatusBadRequest, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Product not found") {
		t.Error("Expected response : Product not found")
	}
}

func TestDeleteNonExistantProduct(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/products/4", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewProductDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": uuid.NewString(),
	}
	request = mux.SetURLVars(request, vars)

	productHandler.Delete(response, request)
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got : %d", http.StatusNotFound, response.Code)
	}
	if !strings.Contains(response.Body.String(), "Product not found") {
		t.Error("Expected response : Product not found")
	}
}

func TestAddProduct(t *testing.T) {
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
	request = request.WithContext(ctx)

	productHandler := NewProductsHandler(NewProductDB())
	productHandler.AddProduct(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestUpdateProduct(t *testing.T) {
	// Creating request body
	body := &data.Product{
		ID:          "a2181017-5c53-422b-b6bc-036b27c04fc8",
		Name:        "addName",
		Description: "addDescription",
		Price:       1,
		SKU:         "abc-abc-abcd",
	}

	request := httptest.NewRequest(http.MethodPut, "/products", nil)
	response := httptest.NewRecorder()

	// Add the body to the context since we arent passing through middleware
	ctx := context.WithValue(request.Context(), KeyProduct{}, body)
	request = request.WithContext(ctx)

	productHandler := NewProductsHandler(NewProductDB())
	productHandler.UpdateProducts(response, request)

	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteExistingProduct(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewProductDB())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}
