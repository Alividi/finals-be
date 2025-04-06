package auth

import (
	"context"
	"net/http"
	"strings"
)

type UserContextKey string
type JWTContextKey string

const (
	UserContext UserContextKey = "USER_CONTEXT_KEY"
	JWTContext  JWTContextKey  = "USER_CONTEXT_KEY"
)

func SetUserContext(r *http.Request, user *User) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, UserContext, user)
	ctx = context.WithValue(ctx, JWTContext, r.Header.Get("Authorization"))
	return r.WithContext(ctx)
}

func GetUserContext(ctx context.Context) *User {
	return ctx.Value(UserContext).(*User)
}

func GetOptionalUserContext(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(UserContext).(*User)
	return user, ok
}

func GetJWTContext(ctx context.Context) string {
	return ctx.Value(JWTContext).(string)
}

func ExtractAuthToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	return strings.TrimPrefix(bearerToken, "Bearer ")
}
