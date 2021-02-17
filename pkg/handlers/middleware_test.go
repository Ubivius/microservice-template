package handlers

import (
	"net/http"
)

// Creating simple http handler to pass to middleware
type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusNoContent)
}

// func TestValidationMiddleware(t *testing.T) {
// 	// Creating handler to pass to middleware validator
// 	handler := &HttpHandler{}
// 	// Creating request body
// 	body := &data.Product{
// 		Name:        "addName",
// 		Description: "addDescription",
// 		Price:       1,
// 		SKU:         "abc-abc-abcd",
// 	}

// 	bodyBytes, _ := json.Marshal(body)

// 	request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(bodyBytes)))
// 	response := httptest.NewRecorder()

// 	productHandler := NewProductsHandler(NewTestLogger())
// 	productHandler.MiddlewareProductValidation(handler)

// 	if response.Code != http.StatusNoContent {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusNoContent, response.Code)
// 	}
// }
