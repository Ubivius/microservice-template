package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

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

func Teapot(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusTeapot)
}

func TestTeapotHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
}

func TestAbs(t *testing.T) {
	got := 1
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
