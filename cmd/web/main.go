package main

import (
	"context"
	"encoding/json"
	"finals-be/app/http"
	"finals-be/internal/config"
	"finals-be/internal/connection"
	"finals-be/internal/util"
	"sync"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

	awsConfig := getAwsConfig(cfg)

	firebaseClient := initFirebaseClient(cfg)
	s3Client := s3.NewFromConfig(*awsConfig)

	defer db.Close()

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		ctx := context.Background()

		server := http.NewServerOption(http.ServerOption{
			Clients: &util.Clients{
				DB:       db,
				Message:  firebaseClient,
				S3Client: s3Client,
			},
			Config: cfg,
		})

		err = server.Run(ctx, cfg)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to run server")
		}

		wg.Done()
	}()

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

func getAwsConfig(cfg *config.Config) *aws.Config {
	creds := aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     cfg.AWS.AccessKeyID,
			SecretAccessKey: cfg.AWS.SecretAccessKey,
		}, nil
	})

	config, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithCredentialsProvider(creds),
		awsconfig.WithRegion(cfg.AWS.Region),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load AWS config")
	}

	return &config
}
