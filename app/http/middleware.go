package http

import (
	"finals-be/internal/config"
	"net/http"
)

type middleware struct {
	cfg config.Config
}

func NewMiddleware(cfg config.Config) *middleware {
	return &middleware{
		cfg: cfg,
	}
}

func (m *middleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
