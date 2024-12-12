package router

import (
	"github.com/go-chi/chi"
	"github.com/gustavo-villar/go-rate-limiter/limiter"
)

func Init() {
	router := chi.NewRouter()
	rate_limiter := limiter.InitializeRateLimiters()

	InitializeMiddlewares(router, rate_limiter)
	InitializeRoutes(router)
	InitializeServer(router)
}
