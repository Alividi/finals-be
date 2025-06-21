package dto

import (
	"finals-be/app/notification/model"
	"finals-be/internal/lib/helper"
)

type NotificationPaginationRequest struct {
	helper.PaginationRequest
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
}

type GetAllNotificationResponse struct {
	Data []*model.Notification `json:"notifications"`
	Meta helper.Pagination     `json:"meta"`
}
