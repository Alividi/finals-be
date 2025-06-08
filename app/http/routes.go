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
	notificationHandler := handler.NewNotificationHandler(s.notificationService)

	RegisterAuthRoutes(router, privateAPI, authHandler)
	RegisterUserRoutes(router, privateAPI, userHandler)
	RegisterProductRoutes(router, privateAPI, productHandler)
	RegisterNotificationRoutes(privateAPI, notificationHandler)

	return router
}

func RegisterAuthRoutes(publicAPI *mux.Router, privateAPI *mux.Router, h *handler.AuthHandler) {
	publicAPI.HandleFunc("/login", h.Login).Methods("POST")
	publicAPI.HandleFunc("/refresh-token", h.RefreshToken).Methods("POST")

	privateAPI.HandleFunc("/logout", h.Logout).Methods("POST")
}

func RegisterUserRoutes(publicAPI *mux.Router, privateAPI *mux.Router, h *handler.UserHandler) {
	privateAPI.HandleFunc("/current-user", h.GetCurrentUser).Methods("GET")
	privateAPI.HandleFunc("/technicians", h.GetTechnicians).Methods("GET")
	privateAPI.HandleFunc("/user-status", h.GetUserStatus).Methods("GET")
}

func RegisterProductRoutes(publicAPI *mux.Router, privateAPI *mux.Router, h *handler.ProductHandler) {
	privateAPI.HandleFunc("/products", h.GetProducts).Methods("GET")
	privateAPI.HandleFunc("/products/{id}", h.GetProductById).Methods("GET")
	privateAPI.HandleFunc("/faqs/{id}", h.GetFAQById).Methods("GET")
}

func RegisterNotificationRoutes(privateApi *mux.Router, h *handler.NotificationHandler) {
	privateApi.HandleFunc("/notifications", h.GetNotifications).Methods("GET")
	privateApi.HandleFunc("/notifications/read/{notification_id}", h.MarkAsRead).Methods("POST")
	privateApi.HandleFunc("/notifications/read", h.MarkAllAsRead).Methods("POST")
}
