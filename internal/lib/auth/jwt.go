package auth

import (
	"errors"
	"finals-be/internal/config"
	"finals-be/internal/lib/helper"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewAuthClaims(ID string, username string, role string, issuer string, expiration time.Time) *AuthClaims {
	return &AuthClaims{
		User: &AuthUser{
			UserID:   ID,
			Username: username,
			Role:     role,
		},
		Purpose: "auth",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
}

func NewTokenClaims(purpose string, recipient string, issuer string, expiration *time.Time) *AuthClaims {
	return &AuthClaims{
		Purpose:   purpose,
		Recipient: recipient,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: issuer,
			ExpiresAt: func() *jwt.NumericDate {
				if expiration == nil {
					return nil
				}
				return jwt.NewNumericDate(*expiration)
			}(),
		},
	}
}

func GenerateToken(claims *AuthClaims, cfg *config.JWTConfig) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecretKey))
}

func ValidateTokenFromRequest(r *http.Request, cfg *config.JWTConfig) (*AuthClaims, error) {
	tokenString := ExtractAuthToken(r)
	if tokenString == "" {
		return nil, helper.NewErrUnauthorized("missing token")
	}

	return ValidateToken(cfg, tokenString)
}

func ValidateToken(cfg *config.JWTConfig, tokenString string) (*AuthClaims, error) {
	claims := &AuthClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, helper.NewErrUnauthorized("unexpected signing method")
		}
		return []byte(cfg.JWTSecretKey), nil
	})

	if err != nil || !token.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, helper.NewErrUnauthorized("token expired")
		}

		return nil, helper.NewErrUnauthorized("invalid token")
	}

	return claims, nil
}
