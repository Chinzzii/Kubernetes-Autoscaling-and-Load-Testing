package main

import (
	"log"
	"net/http"

	"github.com/Chinzzii/Kubernetes-Autoscaling-and-Load-Testing/monolithic/handlers"
)

func main() {
	http.HandleFunc("/compute", handlers.ComputeHandler)
	http.HandleFunc("/data", handlers.DataHandler)
	http.HandleFunc("/status", handlers.StatusHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
