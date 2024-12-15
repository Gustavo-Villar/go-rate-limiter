// Package config handles application configuration by loading environment variables
// from a `.env` file or falling back to the system environment.

package config

import (
	"log"     // Provides logging capabilities for error messages.
	"strings" // Utility functions for manipulating strings.

	"github.com/joho/godotenv" // Library for loading environment variables from a `.env` file.
)

// Init initializes the application configuration by loading environment variables.
func Init() {
	// Attempt to load and override environment variables from a `.env` file.
	err := godotenv.Overload()
	if err != nil {
		// Check if the error indicates that the `.env` file is missing.
		if strings.Contains(string(err.Error()), "no such file or directory") {
			log.Printf("error loading .env file. Continuing without it, getting envs from environment...")
		} else {
			// Log a fatal error and terminate the application for other issues.
			log.Fatalf("fail to read configs: %v", err)
			return
		}
	}
}
