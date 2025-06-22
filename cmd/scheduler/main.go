package main

import (
	"context"
	"encoding/json"
	"finals-be/internal/config"
	"finals-be/internal/connection"
	"finals-be/internal/worker"
	"sync"

	firebaseServicePkg "finals-be/app/notification/service"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

func main() {
	cfg := config.LoadConfigByFile("./config", "dev", "yml")
	log.Info().Msg("Config loaded")

	db, err := connection.NewConnectionManager(cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	firebaseClient := initFirebaseClient(cfg)
	if firebaseClient == nil {
		log.Fatal().Msg("Failed to initialize Firebase client")
	}

	firebaseService := firebaseServicePkg.NewFirebaseService(cfg, db, firebaseClient)
	if firebaseService == nil {
		log.Fatal().Msg("Failed to initialize Firebase service")
	}

	defer db.Close()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	cronScheduler := cron.New(cron.WithSeconds())

	// Schedule: every 1 minute
	insertWorker := worker.NewInsertWorker(db.GetQuery())
	cronScheduler.AddFunc("0 * * * * *", insertWorker.Run) // Every minute at 0s

	checkWorker := worker.NewCheckWorker(db.GetQuery(), firebaseClient, firebaseService)
	cronScheduler.AddFunc("5 * * * * *", checkWorker.Run) // Every minute at 10s

	cronScheduler.Start()

	log.Info().Msg("Scheduler started")
	wg.Wait()
}

func initFirebaseClient(cfg *config.Config) *messaging.Client {
	firebaseCreds, err := json.Marshal(cfg.Firebase)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to marshal firebase credentials")
	}

	opt := option.WithCredentialsJSON(firebaseCreds)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create firebase app")
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create firebase messaging client")
	}

	return client
}
