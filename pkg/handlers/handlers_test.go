package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Ubivius/microservice-template/pkg/data"
	"github.com/gorilla/mux"
)

// Move to util package in Sprint 9, should be a testing specific logger
func NewTestLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestGetProducts(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewTestLogger())
	productHandler.GetProducts(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":2") {
		t.Error("Missing elements from expected results")
	}
}

func TestGetProductByID(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.GetProductByID(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code 200 but got : %d", response.Code)
	}
	if !strings.Contains(response.Body.String(), "\"id\":1") {
		t.Error("Missing elements from expected results")
	}
}

func TestDeleteExistingProduct(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "1",
	}
	request = mux.SetURLVars(request, vars)

	productHandler.Delete(response, request)
	if response.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d but got : %d", http.StatusNoContent, response.Code)
	}
}

func TestDeleteNonExistantProduct(t *testing.T) {
	request := httptest.NewRequest(http.MethodDelete, "/products/4", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewTestLogger())

	// Mocking gorilla/mux vars
	vars := map[string]string{
		"id": "4",
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
		Price:       1.00,
		SKU:         "abc-abc-abcd",
	}
	bodyBytes, _ := json.Marshal(body)
	t.Log(string(bodyBytes))
	// reader := strings.NewReader("{\"name\":\"addName\",\"description\":\"addDescription\",\"price\":1,\"sku\":\"abc-abc-abcd\"}")

	request := httptest.NewRequest(http.MethodPost, "/products", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewTestLogger())
	productHandler.AddProduct(response, request)
	t.Log(response.Body.String())
	t.Fail()
}
