package dto

import "finals-be/app/user/model"

type GetCurrentUserResponse struct {
	Role     string        `json:"role"`
	Username string        `json:"username"`
	Phone    string        `json:"phone"`
	Email    string        `json:"email"`
	Address  *model.Alamat `json:"address"`
}

type GetTechniciansResponse struct {
	ID     string `json:"id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	NoTelp string `json:"no_telp"`
	Status string `json:"status"`
	Base   string `json:"base"`
}
