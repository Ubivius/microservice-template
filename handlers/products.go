package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object inside context
type KeyProduct struct{}

// Product handler used for getting and updating products
type ProductsHandler struct {
	logger *log.Logger
}

func NewProductsHandler(logger *log.Logger) *ProductsHandler {
	return &ProductsHandler{logger}
}

// getProductId extracts the product ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getProductId(request *http.Request) int {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
