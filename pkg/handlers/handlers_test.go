package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// Move to util package in Sprint 9
func NewLogger() *log.Logger {
	return log.New(os.Stdout, "Tests", log.LstdFlags)
}

func TestGetProducts(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	response := httptest.NewRecorder()

	productHandler := NewProductsHandler(NewLogger())
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

	productHandler := NewProductsHandler(NewLogger())

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

	productHandler := NewProductsHandler(NewLogger())

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
