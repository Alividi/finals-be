package worker

import (
	"context"
	"finals-be/app/notification/model"
	"finals-be/app/notification/service"
	"finals-be/internal/connection"
	"finals-be/internal/constants"
	"fmt"
	"time"

	"firebase.google.com/go/v4/messaging"
	"github.com/rs/zerolog/log"
)

type CheckWorker struct {
	db              *connection.SingleInstruction
	notificationSvc *service.FirebaseService
}

func NewCheckWorker(db *connection.SingleInstruction, firebaseClient *messaging.Client, notifSvc *service.FirebaseService) *CheckWorker {
	return &CheckWorker{
		db:              db,
		notificationSvc: notifSvc,
	}
}

func (w *CheckWorker) Run() {
	ctx := context.Background()
	log.Info().Msg("Running check_worker...")

	// Step 1: Get all problematic services + customer user_id
	var results []struct {
		ServiceID      int64  `db:"id"`
		NamaService    string `db:"nama_service"`
		NamaPerusahaan string `db:"nama_perusahaan"`
		CustomerUserID int64  `db:"user_id"`
	}

	query := `
			SELECT 
			s.id,
			s.nama_service,
			c.nama_perusahaan,
			c.user_id
		FROM tbl_service s
		JOIN tbl_customer c ON s.customer_id = c.id
		WHERE s.is_problem = true
	`
	log.Info().Msg("Fetching problematic services...")

	if err := w.db.Select(ctx, &results, query); err != nil {
		log.Error().Err(err).Msg("Failed to get problematic services")
		return
	}

	log.Info().Msgf("Found %d problematic services", len(results))
	if len(results) == 0 {
		log.Info().Msg("No problematic services found")
		return
	}

	now := time.Now()

	for _, svc := range results {
		// Step 1: Get all admin user_ids
		var adminUserIDs []int64
		adminQuery := `
		SELECT u.id
		FROM tbl_admin a
		JOIN tbl_users u ON a.user_id = u.id
		WHERE u.role = 'admin'
	`
		if err := w.db.Select(ctx, &adminUserIDs, adminQuery); err != nil {
			log.Error().Err(err).Msgf("Failed to get admin user IDs for customer %s", svc.NamaPerusahaan)
			continue
		}

		// Step 2: Notify Customer
		customerNotif := model.Notification{
			UserID:    svc.CustomerUserID,
			Judul:     "Gangguan Layanan",
			Deskripsi: fmt.Sprintf("Layanan \"%s\" Anda mengalami gangguan. Silahkan cek cara mengatasinya pada halaman troubleshoot", svc.NamaService),
			Type:      constants.NOTIFICATION_TYPE_SERVICE,
			IsRead:    false,
			CreatedAt: now,
			UpdatedAt: now,
		}
		err := w.notificationSvc.SendNotifications(ctx, customerNotif)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to notify customer user %d", svc.CustomerUserID)
		}
		log.Info().Msgf("Notification sent to customer %s for service %s", svc.NamaPerusahaan, svc.NamaService)

		// Step 3: Notify Admins
		for _, adminUID := range adminUserIDs {
			adminNotif := model.Notification{
				UserID:    adminUID,
				Judul:     "Layanan Customer Bermasalah",
				Deskripsi: fmt.Sprintf("Layanan \"%s\" milik \"%s\" mengalami gangguan.", svc.NamaService, svc.NamaPerusahaan),
				Type:      constants.NOTIFICATION_TYPE_SERVICE,
				IsRead:    false,
				CreatedAt: now,
				UpdatedAt: now,
			}
			err := w.notificationSvc.SendNotifications(ctx, adminNotif)
			if err != nil {
				log.Error().Err(err).Msgf("Failed to notify admin user %d", adminUID)
			}
			log.Info().Msgf("Notification sent to admin user %d for service %s", adminUID, svc.NamaService)
		}
	}

}
