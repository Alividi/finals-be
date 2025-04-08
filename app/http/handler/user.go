package handler

import (
	userService "finals-be/app/user/service"
	"finals-be/internal/lib/auth"
	"finals-be/internal/lib/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService *userService.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService *userService.UserService, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validate,
	}
}

func (u *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserContext(r.Context())

	response, err := u.userService.GetCurrentUser(r.Context(), user.ID)
	if err != nil {
		helper.WriteResponse(r.Context(), w, err, nil)
		return
	}

	helper.WriteResponse(r.Context(), w, err, response)
}
