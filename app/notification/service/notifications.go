package service

import (
	"context"
	"finals-be/app/notification/dto"
	"finals-be/app/notification/repository"
	"finals-be/internal/config"
	"finals-be/internal/connection"

	"github.com/rs/zerolog/log"
)

type NotificationService struct {
	cfg              *config.Config
	db               *connection.SingleInstruction
	notificationRepo repository.INotificationRepository
}

func NewNotificationService(cfg *config.Config, conn *connection.SQLServerConnectionManager) *NotificationService {
	db := conn.GetQuery()
	return &NotificationService{
		cfg:              cfg,
		db:               db,
		notificationRepo: repository.NewNotificationRepository(db),
	}
}

func (s *NotificationService) GetNotifications(ctx context.Context, req dto.NotificationPaginationRequest) (response *dto.GetAllNotificationResponse, err error) {
	log := log.Ctx(ctx).With().Str("service", "GetNotifications").Logger()

	notifications, meta, err := s.notificationRepo.GetNotifications(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get notifications")
		return nil, err
	}

	if len(notifications) == 0 {
		return &dto.GetAllNotificationResponse{}, nil
	}

	return &dto.GetAllNotificationResponse{
		Data: notifications,
		Meta: meta,
	}, nil
}

func (s *NotificationService) MarkAsRead(ctx context.Context, userId int64, notificationId int64) error {
	log := log.Ctx(ctx).With().Str("service", "MarkAsRead").Logger()

	err := s.notificationRepo.MarkAsRead(ctx, userId, notificationId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to mark notification as read")
		return err
	}

	return nil
}

func (s *NotificationService) MarkAllAsRead(ctx context.Context, userId int64) error {
	log := log.Ctx(ctx).With().Str("service", "MarkAllAsRead").Logger()

	err := s.notificationRepo.MarkAllAsRead(ctx, userId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to mark all notifications as read")
		return err
	}

	return nil
}
