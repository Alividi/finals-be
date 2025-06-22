package repository

import (
	"context"
	"finals-be/app/notification/dto"
	"finals-be/app/notification/model"
	"finals-be/internal/connection"
	"finals-be/internal/constants"
	"finals-be/internal/lib/helper"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

type INotificationRepository interface {
	GetFCMTokens(ctx context.Context, userId int64) (tokens []string, err error)
	StoreFCMToken(ctx context.Context, userId int64, token string) error
	DeleteFCMToken(ctx context.Context, userId int64, token string) error
	InsertNotifications(ctx context.Context, req model.Notification) error
	GetNotifications(ctx context.Context, request dto.NotificationPaginationRequest) (notifications []*model.Notification, meta helper.Pagination, err error)
	MarkAsRead(ctx context.Context, userId int64, notificationId int64) error
	MarkAllAsRead(ctx context.Context, userId int64) error
}

type NotificationRepo struct {
	db connection.Connection
}

func NewNotificationRepository(db connection.Connection) *NotificationRepo {
	return &NotificationRepo{db: db}
}

func (r *NotificationRepo) GetFCMTokens(ctx context.Context, userId int64) (tokens []string, err error) {
	query := fmt.Sprintf(`
		SELECT fcm_token
		FROM %s
		WHERE user_id = $1
	`, constants.TABLE_TOKEN)

	tokens = make([]string, 0)
	err = r.db.Select(ctx, &tokens, query, userId)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *NotificationRepo) StoreFCMToken(ctx context.Context, userId int64, token string) error {
	query := fmt.Sprintf(`
		WITH inserted AS (
			INSERT INTO %s (user_id, fcm_token, created_at)
			VALUES ($1, $2, NOW())
		)
		DELETE FROM %s 
		WHERE id IN (
			SELECT id FROM %s 
			WHERE user_id = $1
			ORDER BY created_at DESC 
			OFFSET 5
		);
	`, constants.TABLE_TOKEN, constants.TABLE_TOKEN, constants.TABLE_TOKEN)

	_, err := r.db.Exec(ctx, query, userId, token)
	if err != nil {
		return err
	}

	return nil
}

func (r *NotificationRepo) DeleteFCMToken(ctx context.Context, userId int64, token string) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id = $1 AND fcm_token = $2
	`, constants.TABLE_TOKEN)

	_, err := r.db.Exec(ctx, query, userId, token)
	if err != nil {
		return err
	}

	return nil
}

func (r *NotificationRepo) InsertNotifications(ctx context.Context, req model.Notification) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (user_id, judul, deskripsi, is_read, type, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, constants.TABLE_NOTIFIKASI)

	_, err := r.db.Exec(ctx, query, req.UserID, req.Judul, req.Deskripsi, false, req.Type, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		log.Error().Msgf("Failed to insert notification: %v", err)
		return err
	}

	return nil
}

func (r *NotificationRepo) GetNotifications(ctx context.Context, request dto.NotificationPaginationRequest) (notifications []*model.Notification, meta helper.Pagination, err error) {
	notifications = make([]*model.Notification, 0)
	offset := (request.Page - 1) * request.PageSize
	meta = helper.Pagination{}

	var queryBuilder strings.Builder
	queryBuilder.WriteString(fmt.Sprintf(`
		SELECT id, user_id, judul, deskripsi, is_read, type, created_at, updated_at
		FROM %s
		WHERE user_id = $1
	`, constants.TABLE_NOTIFIKASI))

	var args []interface{}
	args = append(args, request.UserID)
	argIndex := 2

	if request.Type != "" {
		queryBuilder.WriteString(fmt.Sprintf(" AND type = $%d", argIndex))
		args = append(args, request.Type)
		argIndex++
	}

	queryBuilder.WriteString(fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1))
	args = append(args, request.PageSize, offset)

	query := queryBuilder.String()

	err = r.db.Select(ctx, &notifications, query, args...)
	if err != nil {
		return nil, meta, err
	}

	var countQuery strings.Builder
	countQuery.WriteString(fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s
		WHERE user_id = $1
	`, constants.TABLE_NOTIFIKASI))

	var countArgs []interface{}
	countArgs = append(countArgs, request.UserID)
	countIndex := 2

	if request.Type != "" {
		countQuery.WriteString(fmt.Sprintf(" AND type = $%d", countIndex))
		countArgs = append(countArgs, request.Type)
	}

	var total int
	err = r.db.Get(ctx, &total, countQuery.String(), countArgs...)
	if err != nil {
		return nil, meta, err
	}

	meta = helper.NewPagination(request.Page, request.PageSize, total)

	return notifications, meta, nil
}

func (r *NotificationRepo) MarkAsRead(ctx context.Context, userId int64, notificationId int64) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET is_read = true
		WHERE id = $1 AND user_id = $2
	`, constants.TABLE_NOTIFIKASI)

	_, err := r.db.Exec(ctx, query, notificationId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *NotificationRepo) MarkAllAsRead(ctx context.Context, userId int64) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET is_read = true
		WHERE user_id = $1
	`, constants.TABLE_NOTIFIKASI)

	_, err := r.db.Exec(ctx, query, userId)
	if err != nil {
		return err
	}

	return nil
}
