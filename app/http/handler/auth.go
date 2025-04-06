package handler

import (
	"finals-be/app/auth/dto"
	auth "finals-be/app/auth/service"
	"finals-be/internal/lib/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService *auth.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(authService *auth.AuthService, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validate,
	}
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	payload := dto.LoginRequest{}
	err := helper.ReadRequest(r, &payload)
	if err != nil {
		helper.WriteResponse(r.Context(), w, err, nil)
		return
	}

	err = a.validate.Struct(payload)
	if err != nil {
		helper.WriteResponse(r.Context(), w, helper.NewErrBadRequest(err.Error()), nil)
		return
	}

	response, err := a.authService.Login(r.Context(), payload)
	if err != nil {
		helper.WriteResponse(r.Context(), w, err, nil)
		return
	}

	helper.WriteResponse(r.Context(), w, err, response)
}
