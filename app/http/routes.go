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

	privateAPI := router.PathPrefix("").Subrouter()
	privateAPI.Use(middleware.Authorization)

	authHandler := handler.NewAuthHandler(s.authService, s.validate)

	RegisterAuthRoutes(router, privateAPI, authHandler)

	return router
}

func RegisterAuthRoutes(publicAPI *mux.Router, privateAPI *mux.Router, h *handler.AuthHandler) {
	publicAPI.HandleFunc("/login", h.Login).Methods("POST")
	publicAPI.HandleFunc("/refresh-token", h.RefreshToken).Methods("POST")

	privateAPI.HandleFunc("/logout", h.Logout).Methods("POST")
}
