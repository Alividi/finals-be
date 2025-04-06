package http

import (
	"context"
	"finals-be/app/http/handler"
	"finals-be/internal/config"

	"github.com/gorilla/mux"
)

func RegisterRoutes(ctx context.Context, s *Server, cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	middleware := NewMiddleware(*cfg)

	authHandler := handler.NewAuthHandler(s.authService, s.validate)

	router.Use(middleware.Authorization)

	RegisterAuthRoutes(router, authHandler)

	return router
}

func RegisterAuthRoutes(api *mux.Router, h *handler.AuthHandler) {
	api.HandleFunc("/login", h.Login).Methods("POST") // handler
}
