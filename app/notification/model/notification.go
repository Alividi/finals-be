package model

type Notification struct {
	ID        int64  `json:"id" db:"id"`
	UserID    int64  `json:"user_id" db:"user_id"`
	IsRead    bool   `json:"is_read" db:"is_read"`
	Judul     string `json:"judul" db:"judul"`
	Type      string `json:"type" db:"type"`
	Deskripsi string `json:"deskripsi" db:"deskripsi"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
