package handler

import (
	dtonotification "finals-be/app/notification/dto"
	notification "finals-be/app/notification/service"
	"finals-be/internal/lib/auth"
	"finals-be/internal/lib/helper"
	"net/http"
)

type NotificationHandler struct {
	notificationService *notification.NotificationService
}

func NewNotificationHandler(notificationService *notification.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

func (h *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request dtonotification.NotificationPaginationRequest

	request.Page = helper.GetQueryInt(r, "page", 1)
	request.PageSize = helper.GetQueryInt(r, "page_size", 20)
	request.Type = helper.GetQueryString(r, "type", "")

	userCtx := auth.GetUserContext(ctx)
	request.UserID = userCtx.ID

	notifications, err := h.notificationService.GetNotifications(ctx, request)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, notifications)
}

func (h *NotificationHandler) MarkAsRead(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userCtx := auth.GetUserContext(ctx)
	notificationID := helper.GetURLParamInt64(r, "notification_id")

	err := h.notificationService.MarkAsRead(ctx, userCtx.ID, notificationID)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, nil)
}

func (h *NotificationHandler) MarkAllAsRead(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userCtx := auth.GetUserContext(ctx)

	err := h.notificationService.MarkAllAsRead(ctx, userCtx.ID)
	if err != nil {
		helper.WriteResponse(ctx, w, err, nil)
		return
	}

	helper.WriteResponse(ctx, w, nil, nil)
}
