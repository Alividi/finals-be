package dto

import "finals-be/app/user/model"

type GetCurrentUserResponse struct {
	Role     string        `json:"role"`
	Username string        `json:"username"`
	Phone    string        `json:"phone"`
	Email    string        `json:"email"`
	Address  *model.Alamat `json:"address"`
}
