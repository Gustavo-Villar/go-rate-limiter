// Package limiter implements the logic for rate limiting requests based on IP or token.
// It includes middleware integration and uses Redis as the storage mechanism.

package limiter

import (
	"context"  // Provides context for Redis operations.
	"fmt"      // Provides formatted I/O functions.
	"log"      // Provides logging utilities.
	"net"      // Handles network-related operations.
	"net/http" // Implements HTTP client and server.
	"os"       // Handles environment variables.
	"strconv"  // Converts strings to integers.
	"strings"  // Provides string manipulation functions.
	"time"     // Handles time-related functions.

	"github.com/go-redis/redis/v8" // Redis client library.
)

// RateLimiter contains configurations and methods for handling rate limiting.
type RateLimiter struct {
	store          Store         // Storage mechanism for rate limit data.
	rateLimitIP    int           // Max requests per second per IP.
	rateLimitToken int           // Max requests per second per token.
	blockDuration  time.Duration // Duration for which requests are blocked after exceeding the limit.
}

// NewRateLimiter initializes and returns a new RateLimiter instance.
func NewRateLimiter(store Store) *RateLimiter {
	// Retrieve rate limit values and block duration from environment variables.
	rateLimitIP, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_IP"))
	rateLimitToken, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	blockDuration, _ := strconv.Atoi(os.Getenv("BLOCK_DURATION"))

	// Create and return a configured RateLimiter instance.
	return &RateLimiter{
		store:          store,
		rateLimitIP:    rateLimitIP,
		rateLimitToken: rateLimitToken,
		blockDuration:  time.Duration(blockDuration) * time.Second,
	}
}

// getIP extracts the client's IP address from the HTTP request.
// It first checks common headers (e.g., "X-Forwarded-For") and then falls back to the remote address.
func getIP(r *http.Request) string {
	headers := []string{"X-Forwarded-For", "X-Real-IP"}
	for _, header := range headers {
		ip := r.Header.Get(header)
		if ip != "" {
			return strings.Split(ip, ",")[0]
		}
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

// CheckRateLimit verifies if the request is allowed based on the configured rate limit for IP or token.
func (rl *RateLimiter) CheckRateLimit(ip, token string) (bool, error) {
	var limit int
	// Determine which limit to use based on whether a token is provided.
	if token != "" {
		limit = rl.rateLimitToken
	} else {
		limit = rl.rateLimitIP
	}

	// Use the token as the key if available; otherwise, use the IP.
	key := ip
	if token != "" {
		key = token
	}

	// Check if the request is allowed using the store's Allow method.
	return rl.store.Allow(key, limit, rl.blockDuration)
}

// Middleware creates an HTTP middleware that applies the rate limiter to incoming requests.
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract client IP and token from the request.
		ip := getIP(r)
		token := r.Header.Get("API_KEY")

		// Check if the request is allowed based on rate limiting rules.
		allowed, err := rl.CheckRateLimit(ip, token)
		if err != nil {
			// Respond with an internal server error if rate limiting fails.
			fmt.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Respond with "Too Many Requests" if the rate limit is exceeded.
		if !allowed {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}

		// Forward the request to the next handler if allowed.
		next.ServeHTTP(w, r)
	})
}

// InitializeRateLimiters sets up the Redis client and returns a configured RateLimiter instance.
func InitializeRateLimiters() *RateLimiter {
	// Create a new Redis client using configuration from environment variables.
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	})

	// Check if the Redis connection is successful.
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("could not connect to Redis: %v", err)
		return nil
	}

	// Create a Redis-backed store and initialize the RateLimiter.
	redis_store := NewRedisStore(rdb)
	return NewRateLimiter(redis_store)
}
