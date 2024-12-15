// Entry point of the Go Rate Limiter application.
// This file initializes the application configuration and sets up the HTTP router.

package main

import (
	"github.com/gustavo-villar/go-rate-limiter/config" // Package for application configuration initialization.
	"github.com/gustavo-villar/go-rate-limiter/router" // Package for setting up the HTTP router.
)

func main() {
	// Initialize application configuration, including environment variables and dependencies.
	config.Init()

	// Initialize the HTTP router and start the server on the configured port.
	router.Init()
}
