package router

import (
	"github.com/go-chi/chi"
	"github.com/gustavo-villar/go-rate-limiter/handler"
)

func InitializeRoutes(router *chi.Mux) {
	router.Get("/api/v1/healthz", handler.HealthzHandler)
}
