package handlers

import (
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/database"
	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object inside context
type KeyProduct struct{}

// ProductsHandler contains the items common to all product handler functions
type ProductsHandler struct {
	db database.ProductDB
}

// NewProductsHandler returns a pointer to a ProductsHandler with the logger passed as a parameter
func NewProductsHandler(db database.ProductDB) *ProductsHandler {
	return &ProductsHandler{db}
}

// getProductID extracts the product ID from the URL
// The verification of this variable is handled by gorilla/mux
// We panic if it is not valid because that means gorilla is failing
func getProductID(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]

	return id
}
