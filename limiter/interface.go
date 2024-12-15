// Package limiter defines the interface for rate-limiting functionality,
// allowing for flexibility in the underlying storage mechanism.

package limiter

import "time" // Provides time-related functions and types.

// Store is an interface that defines the contract for implementing
// rate-limiting storage and logic.
type Store interface {
	// Allow checks if a request associated with the given key is allowed
	// based on the specified request limit and duration.
	//
	// Parameters:
	// - key: A unique identifier for the entity being rate-limited (e.g., IP or token).
	// - limit: The maximum number of requests allowed within the duration.
	// - duration: The time window during which the limit applies.
	//
	// Returns:
	// - bool: True if the request is allowed, false otherwise.
	// - error: An error object if something goes wrong during the process.
	Allow(key string, limit int, duration time.Duration) (bool, error)
}
