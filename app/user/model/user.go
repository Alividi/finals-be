package model

type User struct {
	ID           string  `db:"id"`
	Username     string  `db:"username"`
	Password     string  `db:"password"`
	Nama         string  `db:"nama"`
	Email        string  `db:"email"`
	NoTelp       string  `db:"no_telp"`
	Role         string  `db:"role"`
	RefreshToken *string `db:"refresh_token"`
}

type CustomerDetail struct {
	User
	CustomerWithAlamat
}

type TeknisiDetail struct {
	User
}

type AdminDetail struct {
	User
}
