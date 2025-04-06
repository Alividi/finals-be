package main

import (
	"context"
	"finals-be/app/http"
	"finals-be/internal/config"
	"finals-be/internal/connection"
	"log"
	"time"
)

func main() {

	cfg := config.Config{
		App: config.AppConfig{
			Name:     "finals-app",
			HttpPort: 8080,
			AppURL:   "http://localhost:8080",
		},
		DB: config.DatabaseConfig{
			Driver:              "postgres",
			URL:                 "postgres://postgres:root@localhost:5432/finals_db?sslmode=disable",
			MaxIdleConnections:  20,
			MaxOpenConnections:  50,
			MaxIdleDuration:     60,
			MaxLifeTimeDuration: 300,
			MigrationPath:       "file://migrations",
		},
		JWT: config.JWTConfig{
			JWTSecretKey:              "anjay",
			LoginExpirationDuration:   24 * time.Hour,
			RefreshExpirationDuration: 7 * 24 * time.Hour,
		},
	}

	db, err := connection.NewConnectionManager(cfg.DB)
	if err != nil {
		log.Fatal("Failed to connect to databases")
	}

	ctx := context.Background()

	server := http.NewServerOption(http.ServerOption{
		Config: &cfg,
		DB:     db,
	})

	err = server.Run(ctx, &cfg)
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
