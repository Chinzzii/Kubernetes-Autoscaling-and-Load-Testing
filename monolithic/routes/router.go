package routes

import (
	// "net/http"
	"github.com/Chinzzii/Kubernetes-Autoscaling-and-Load-Testing/handlers"

	"github.com/gorilla/mux"
)

// InitRouter initializes and returns the application router.
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/status", handlers.StatusHandler).Methods("GET")
	router.HandleFunc("/data", handlers.DataHandler).Methods("GET")
	router.HandleFunc("/compute", handlers.ComputeHandler).Methods("GET")

	// Add middleware for logging

	return router
}
