package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/Chinzzii/Kubernetes-Autoscaling-and-Load-Testing/monolithic/utils"
)

// ComputeHandler simulates heavy computation
func ComputeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Simulate heavy computation
	result := 0.0
	for i := 1; i <= 1000000; i++ {
		result += math.Sqrt(float64(i))
	}

	elapsed := time.Since(start)
	response := map[string]interface{}{
		"result":  result,
		"elapsed": elapsed.String(),
	}
	JSONResponse(w, http.StatusOK, response)
}
