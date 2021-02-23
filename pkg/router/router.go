package router

import (
	"log"
	"net/http"

	"github.com/Ubivius/microservice-template/pkg/handlers"
	"github.com/gorilla/mux"
)

// Mux route handling with gorilla/mux
func New(productHandler *handlers.ProductsHandler, logger *log.Logger) *mux.Router {
	router := mux.NewRouter()

	// Get Router
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.GetProductByID)

	// Put router
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", productHandler.UpdateProducts)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	// Post router
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	// Delete router
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete)

	return router
}
