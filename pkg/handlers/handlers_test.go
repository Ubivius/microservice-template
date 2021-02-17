package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
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
	if !strings.Contains(response.Body.String(), "id") {
		t.Error("Missing elements from expected results")
	}
}
