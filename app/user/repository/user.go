package repository

import (
	"context"
	"database/sql"
	"finals-be/app/user/model"
	"finals-be/internal/connection"
	"finals-be/internal/constants"
	"finals-be/internal/lib/helper"
	"fmt"
)

type IUserRepository interface {
	GetByUsername(ctx context.Context, username string) (usr *model.User, err error)
	StoreRefreshToken(ctx context.Context, username string, refreshToken *string) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (usr *model.User, err error)
	GetUserDetail(ctx context.Context, userId string) (usr *model.User, err error)
	GetCustomerDetail(ctx context.Context, userId string) (usr *model.CustomerDetail, err error)
	GetAlamatUser(ctx context.Context, userId string) (alamat *model.Alamat, err error)
	GetTechnicians(ctx context.Context) ([]*model.Teknisi, error)
}

type UserRepository struct {
	db connection.Connection
}

func NewUserRepository(db connection.Connection) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (usr *model.User, err error) {

	usr = &model.User{}

	query := fmt.Sprintf(
		`SELECT 
			id, 
			username, 
			password, 
			role, 
			refresh_token 
		FROM %s 
		WHERE username = $1`,
		constants.TABLE_USERS)

	err = r.db.
		QueryRow(ctx, query, username).
		Scan(&usr.ID, &usr.Username, &usr.Password, &usr.Role, &usr.RefreshToken)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("User not found")
		}
		return nil, err
	}

	return
}

func (r *UserRepository) StoreRefreshToken(ctx context.Context, username string, refreshToken *string) error {
	query := fmt.Sprintf(`
		UPDATE %s 
			SET refresh_token = $1
		WHERE username = $2
	`, constants.TABLE_USERS)

	_, err := r.db.Exec(ctx, query, refreshToken, username)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetByRefreshToken(ctx context.Context, refreshToken string) (usr *model.User, err error) {
	usr = &model.User{}

	query := fmt.Sprintf(`
		SELECT
			id, 
			username, 
			role, 
			refresh_token
		FROM %s
		WHERE refresh_token = $1
	`, constants.TABLE_USERS)

	err = r.db.
		QueryRow(ctx, query, refreshToken).
		Scan(&usr.ID, &usr.Username, &usr.Role, &usr.RefreshToken)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("User not found")
		}
		return nil, err
	}

	return
}

func (r *UserRepository) GetCustomerDetail(ctx context.Context, userId string) (usr *model.CustomerDetail, err error) {
	usr = &model.CustomerDetail{}

	query := fmt.Sprintf(`
		SELECT 
			tbl_users.id, tbl_users.username, tbl_users.role
			tbl_customer.nama_perusahaan, tbl_customer.email_perusahaan, tbl_customer.no_telp_perusahaan, tbl_customer.no_npwp_perusahaan, 
			tbl_alamat.provinsi, tbl_alamat.kabupaten, tbl_alamat.kecamatan, tbl_alamat.kelurahan, tbl_alamat.rt, tbl_alamat.rw, tbl_alamat.alamat, tbl_alamat.latitude, tbl_alamat.longitude
		FROM tbl_users
		INNER JOIN tbl_customer ON tbl_users.id = tbl_customer.user_id
		INNER JOIN tbl_alamat ON tbl_alamat.customer_id = tbl_users.id
		WHERE tbl_users.id = $1
	`)

	err = r.db.
		QueryRow(ctx, query, userId).
		Scan(
			&usr.ID, &usr.Username, &usr.Role,
			&usr.Customer.NamaPerusahaan, &usr.Customer.EmailPerusahaan, &usr.Customer.NoTelpPerusahaan, &usr.Customer.NoNpwpPerusahaan,
			&usr.Alamat.Provinsi, &usr.Alamat.Kabupaten, &usr.Alamat.Kecamatan, &usr.Alamat.Kelurahan, &usr.Alamat.RT, &usr.Alamat.RW, &usr.Alamat.Alamat, &usr.Alamat.Latitude, &usr.Alamat.Longitude,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("User not found")
		}
		return nil, err
	}

	return
}

func (r *UserRepository) GetUserDetail(ctx context.Context, userId string) (usr *model.User, err error) {
	usr = &model.User{}

	query := fmt.Sprintf(`
		SELECT 
			id, username, role, no_telp, email
		FROM %s
		WHERE id = $1
	`, constants.TABLE_USERS)

	err = r.db.
		QueryRow(ctx, query, userId).
		Scan(&usr.ID, &usr.Username, &usr.Role, &usr.NoTelp, &usr.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("User not found")
		}
		return nil, err
	}

	return
}

func (r *UserRepository) GetAlamatUser(ctx context.Context, userId string) (alamat *model.Alamat, err error) {
	alamat = &model.Alamat{}

	query := `
		SELECT 
			provinsi, kabupaten, kecamatan, kelurahan, rt, rw, alamat, latitude, longitude  
		FROM tbl_alamat
		INNER JOIN tbl_customer ON tbl_customer.id = tbl_alamat.customer_id
		WHERE tbl_customer.user_id = $1
	`

	err = r.db.
		QueryRow(ctx, query, userId).
		Scan(&alamat.Provinsi, &alamat.Kabupaten, &alamat.Kecamatan, &alamat.Kelurahan, &alamat.RT, &alamat.RW, &alamat.Alamat, &alamat.Latitude, &alamat.Longitude)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("User not found")
		}
		return nil, err
	}

	return
}

func (r *UserRepository) GetTechnicians(ctx context.Context) (technicians []*model.Teknisi, err error) {
	query := fmt.Sprintf(`
		SELECT 
			u.id, u.nama, u.email, u.no_telp, t.status, t.base 
		FROM %s u
		JOIN %s t ON u.id = t.user_id
		WHERE u.role = 'teknisi'
	`, constants.TABLE_USERS, constants.TABLE_TEKNISI)

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		tech := &model.Teknisi{}
		err = rows.Scan(&tech.ID, &tech.Nama, &tech.Email, &tech.NoTelp, &tech.Status, &tech.Base)
		if err != nil {
			return nil, err
		}
		technicians = append(technicians, tech)
	}

	return technicians, nil
}
