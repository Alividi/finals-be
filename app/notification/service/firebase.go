package service

import (
	"context"
	"finals-be/app/notification/model"
	"finals-be/app/notification/repository"
	"finals-be/internal/config"
	"finals-be/internal/connection"

	"firebase.google.com/go/v4/messaging"
	"github.com/rs/zerolog/log"
)

type FirebaseService struct {
	client           *messaging.Client
	cfg              *config.Config
	db               *connection.SingleInstruction
	notificationRepo repository.INotificationRepository
}

func NewFirebaseService(cfg *config.Config, conn *connection.SQLServerConnectionManager, message *messaging.Client) *FirebaseService {
	db := conn.GetQuery()
	return &FirebaseService{
		client:           message,
		cfg:              cfg,
		db:               db,
		notificationRepo: repository.NewNotificationRepository(db),
	}
}

func (s *FirebaseService) SendNotification(ctx context.Context, token string, title string, body string) error {
	log := log.Ctx(ctx).With().Str("service", "SendNotification").Logger()
	msg := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	response, err := s.client.Send(ctx, msg)
	if err != nil {
		return err
	}

	log.Info().Msgf("Successfully sent message: %s", response)
	return nil
}

func (s *FirebaseService) SendNotifications(ctx context.Context, req model.Notification) error {
	log := log.Ctx(ctx).With().Str("service", "SendNotifications").Logger()

	tokens, err := s.notificationRepo.GetFCMTokens(ctx, req.UserID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user notification token")
		return err
	}

	if len(tokens) == 0 {
		return nil
	}

	var messages []*messaging.Message
	for _, token := range tokens {
		msg := &messaging.Message{
			Token: token,
			Notification: &messaging.Notification{
				Title: req.Judul,
				Body:  req.Deskripsi,
			},
		}
		messages = append(messages, msg)
	}

	_, err = s.client.SendEach(context.Background(), messages)
	if err != nil {
		return err
	}

	err = s.notificationRepo.InsertNotifications(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert notification")
	}

	return nil
}

func (s *FirebaseService) GetUserNotificationToken(ctx context.Context, userID int64) ([]string, error) {
	log := log.Ctx(ctx).With().Str("service", "GetUserNotificationToken").Logger()

	tokens, err := s.notificationRepo.GetFCMTokens(ctx, userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user notification token")
		return nil, err
	}

	return tokens, nil
}
