package http

import (
	"context"
	"finals-be/internal/config"

	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(ctx context.Context, s *Server, cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	middleware := NewMiddleware(*cfg)

	router.Use(middleware.Authorization)

	RegisterAuthRoutes(router)

	return router
}

func RegisterAuthRoutes(api *mux.Router) {
	api.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mantap"))
	}).Methods("POST") // handler
}
