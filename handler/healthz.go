// Package handler defines HTTP handlers for the application, including the health check endpoint.

package handler

import (
	"encoding/json" // Provides functions for JSON encoding and decoding.
	"net/http"      // Implements HTTP client and server functionalities.
)

// Status represents the health status of the application in JSON format.
type Status struct {
	Status string `json:"status"` // JSON field for the status message.
}

// HealthzHandler handles health check requests to verify the application's status.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	// Create a status object indicating the application is running.
	status := Status{
		Status: "UP",
	}

	// Set the response header to indicate the content type is JSON.
	w.Header().Set("Content-Type", "application/json")

	// Encode the status object as JSON and write it to the response.
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		// If JSON encoding fails, respond with an internal server error.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
