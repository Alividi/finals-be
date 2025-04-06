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
	StoreRefreshToken(ctx context.Context, username string, refreshToken string) error
	GetByRefreshToken(ctx context.Context, refreshToken string) (usr *model.User, err error)
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

func (r *UserRepository) StoreRefreshToken(ctx context.Context, username string, refreshToken string) error {
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
