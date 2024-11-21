package handlers

import (
	"net/http"

	"github.com/Chinzzii/Kubernetes-Autoscaling-and-Load-Testing/monolithic/utils"
)

// DataHandler simulates fetching data
func DataHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "Here is some data",
		"items":   []string{"item1", "item2", "item3"},
	}
	utils.JSONResponse(w, http.StatusOK, data)
}
