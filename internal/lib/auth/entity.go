package auth

import "github.com/golang-jwt/jwt/v5"

type GetCurrentUserResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       User   `json:"data"`
}

type User struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role_id"`
}

type AuthClaims struct {
	User      *AuthUser `json:"user"`
	Purpose   string    `json:"purpose"`
	Recipient string    `json:"recipient"`
	jwt.RegisteredClaims
}

type AuthUser struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
