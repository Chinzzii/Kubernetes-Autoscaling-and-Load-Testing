package handlers

import (
	"net/http"

	"github.com/Chinzzii/Kubernetes-Autoscaling-and-Load-Testing/monolithic/utils"
)

// StatusHandler returns the server status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "Server is running"}
	utils.JSONResponse(w, http.StatusOK, response)
}
