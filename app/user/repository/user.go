package repository

import (
	"context"
	"database/sql"
	"finals-be/app/user/dto"
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
	GetUserDetail(ctx context.Context, userId int64) (usr *model.User, err error)
	GetCustomerDetail(ctx context.Context, userId int64) (usr *model.CustomerDetail, err error)
	GetAlamatUser(ctx context.Context, userId int64) (alamat *model.Alamat, err error)
	GetTechnicians(ctx context.Context) ([]*model.Teknisi, error)
	GetUserStatus(ctx context.Context, userId int64) (*dto.UserStatus, error)
	GetCustomerByCustomerID(ctx context.Context, customerId int64) (*model.CustomerDetail, error)
	GetTechnicianByTeknisiID(ctx context.Context, teknisiId int64) (*model.Teknisi, error)
	GetTechnicianByUserId(ctx context.Context, userId int64) (*model.Teknisi, error)
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

func (r *UserRepository) GetCustomerDetail(ctx context.Context, userId int64) (usr *model.CustomerDetail, err error) {
	usr = &model.CustomerDetail{}

	query := fmt.Sprintf(`
		SELECT 
			u.id, u.username, u.role,
			c.id, c.nama_perusahaan, c.email_perusahaan, c.no_telp_perusahaan, c.no_npwp_perusahaan, 
			a.provinsi, a.kabupaten, a.kecamatan, a.kelurahan, a.rt, a.rw, a.alamat, a.latitude, a.longitude
		FROM %s u
		INNER JOIN %s c ON u.id = c.user_id
		INNER JOIN %s a ON a.customer_id = u.id
		WHERE u.id = $1
	`, constants.TABLE_USERS, constants.TABLE_CUSTOMER, constants.TABLE_ALAMAT)

	err = r.db.
		QueryRow(ctx, query, userId).
		Scan(
			&usr.ID, &usr.Username, &usr.Role,
			&usr.Customer.CustomerID, &usr.Customer.NamaPerusahaan, &usr.Customer.EmailPerusahaan, &usr.Customer.NoTelpPerusahaan, &usr.Customer.NoNpwpPerusahaan,
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

func (r *UserRepository) GetUserDetail(ctx context.Context, userId int64) (usr *model.User, err error) {
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

func (r *UserRepository) GetAlamatUser(ctx context.Context, userId int64) (alamat *model.Alamat, err error) {
	alamat = &model.Alamat{}

	query := fmt.Sprintf(`
		SELECT 
			provinsi, kabupaten, kecamatan, kelurahan, rt, rw, alamat, latitude, longitude  
		FROM %s a
		INNER JOIN %s c ON c.id = a.customer_id
		WHERE c.user_id = $1
	`, constants.TABLE_ALAMAT, constants.TABLE_CUSTOMER)

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
			t.id, u.nama, u.email, u.no_telp, t.status, t.base 
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

func (r *UserRepository) GetUserStatus(ctx context.Context, userId int64) (*dto.UserStatus, error) {
	query := fmt.Sprintf(`
		SELECT 
			COUNT(*) AS notification_count 
		FROM %s 
		WHERE user_id = $1 AND is_read = false
	`, constants.TABLE_NOTIFIKASI)

	var notificationCount int
	err := r.db.QueryRow(ctx, query, userId).Scan(&notificationCount)
	if err != nil {
		return nil, err
	}

	return &dto.UserStatus{
		NotificationCount: notificationCount,
	}, nil
}

func (r *UserRepository) GetCustomerByCustomerID(ctx context.Context, customerId int64) (*model.CustomerDetail, error) {
	customer := &model.CustomerDetail{}

	query := fmt.Sprintf(`
		SELECT 
			u.id, u.username, u.role,
			c.id, c.nama_perusahaan, c.email_perusahaan, c.no_telp_perusahaan, c.no_npwp_perusahaan, 
			a.provinsi, a.kabupaten, a.kecamatan, a.kelurahan, a.rt, a.rw, a.alamat, a.latitude, a.longitude
		FROM %s u
		INNER JOIN %s c ON u.id = c.user_id
		INNER JOIN %s a ON a.customer_id = u.id
		WHERE c.id = $1
	`, constants.TABLE_USERS, constants.TABLE_CUSTOMER, constants.TABLE_ALAMAT)

	err := r.db.
		QueryRow(ctx, query, customerId).
		Scan(
			&customer.ID, &customer.Username, &customer.Role,
			&customer.Customer.CustomerID, &customer.Customer.NamaPerusahaan, &customer.Customer.EmailPerusahaan, &customer.Customer.NoTelpPerusahaan, &customer.Customer.NoNpwpPerusahaan,
			&customer.Alamat.Provinsi, &customer.Alamat.Kabupaten, &customer.Alamat.Kecamatan, &customer.Alamat.Kelurahan, &customer.Alamat.RT, &customer.Alamat.RW, &customer.Alamat.Alamat, &customer.Alamat.Latitude, &customer.Alamat.Longitude,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("Customer not found")
		}
		return nil, err
	}

	return customer, nil
}

func (r *UserRepository) GetTechnicianByTeknisiID(ctx context.Context, teknisiId int64) (*model.Teknisi, error) {
	teknisi := &model.Teknisi{}

	query := fmt.Sprintf(`
		SELECT 
			t.id, u.nama, u.email, u.no_telp, t.status, t.base 
		FROM %s u
		JOIN %s t ON u.id = t.user_id
		WHERE u.role = 'teknisi' AND t.id = $1
	`, constants.TABLE_USERS, constants.TABLE_TEKNISI)

	err := r.db.QueryRow(ctx, query, teknisiId).Scan(&teknisi.ID, &teknisi.Nama, &teknisi.Email, &teknisi.NoTelp, &teknisi.Status, &teknisi.Base)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("Technician not found")
		}
		return nil, err
	}

	return teknisi, nil
}

func (r *UserRepository) GetTechnicianByUserId(ctx context.Context, userId int64) (*model.Teknisi, error) {
	teknisi := &model.Teknisi{}

	query := fmt.Sprintf(`
		SELECT 
			t.id, u.nama, u.email, u.no_telp, t.status, t.base 
		FROM %s u
		JOIN %s t ON u.id = t.user_id
		WHERE u.role = 'teknisi' AND u.id = $1
	`, constants.TABLE_USERS, constants.TABLE_TEKNISI)

	err := r.db.QueryRow(ctx, query, userId).Scan(&teknisi.ID, &teknisi.Nama, &teknisi.Email, &teknisi.NoTelp, &teknisi.Status, &teknisi.Base)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewErrNotFound("Technician not found")
		}
		return nil, err
	}

	return teknisi, nil
}
