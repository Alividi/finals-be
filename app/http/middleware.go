package http

import (
	"finals-be/internal/config"
	"finals-be/internal/lib/auth"
	"finals-be/internal/lib/helper"
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
		claims, err := auth.ValidateTokenFromRequest(r, &m.cfg.JWT)
		if err != nil {
			helper.WriteResponse(r.Context(), w, err, nil)
			return
		}

		if claims == nil || claims.User == nil {
			helper.WriteResponse(r.Context(), w, helper.NewErrUnauthorized("Invalid token"), nil)
			return
		}

		user := &auth.User{
			ID:       claims.User.UserID,
			Username: claims.User.Username,
			Role:     claims.User.Role,
		}

		r = auth.SetUserContext(r, user)

		next.ServeHTTP(w, r)
	})
}
