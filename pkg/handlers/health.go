package handlers

import (
	"net/http"
)

// LivenessCheck determine when the application needs to be restarted
func (productHandler *ProductsHandler) LivenessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
}

//ReadinessCheck verifies that the application is ready to accept requests
func (productHandler *ProductsHandler) ReadinessCheck(responseWriter http.ResponseWriter, request *http.Request) {
	err := productHandler.db.PingDB()
	if err != nil {
		log.Error(err, "DB unavailable")
		http.Error(responseWriter, "DB unavailable", http.StatusServiceUnavailable)
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
}
