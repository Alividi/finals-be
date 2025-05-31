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
	userHandler := handler.NewUserHandler(s.userService, s.validate)
	productHandler := handler.NewProductHandler(s.productService)

	RegisterAuthRoutes(router, privateAPI, authHandler)
	RegisterUserRoutes(router, privateAPI, userHandler)
	RegisterProductRoutes(router, privateAPI, productHandler)

	return router
}

func RegisterAuthRoutes(publicAPI *mux.Router, privateAPI *mux.Router, h *handler.AuthHandler) {
	publicAPI.HandleFunc("/login", h.Login).Methods("POST")
	publicAPI.HandleFunc("/refresh-token", h.RefreshToken).Methods("POST")

	privateAPI.HandleFunc("/logout", h.Logout).Methods("POST")
}

func RegisterUserRoutes(publicAPI *mux.Router, privateAPI *mux.Router, h *handler.UserHandler) {
	privateAPI.HandleFunc("/current-user", h.GetCurrentUser).Methods("GET")
}

func RegisterProductRoutes(publicAPI *mux.Router, privateAPI *mux.Router, h *handler.ProductHandler) {
	privateAPI.HandleFunc("/products", h.GetProducts).Methods("GET")
	privateAPI.HandleFunc("/products/{id}", h.GetProductById).Methods("GET")
	privateAPI.HandleFunc("/faqs/{id}", h.GetFAQById).Methods("GET")
}
