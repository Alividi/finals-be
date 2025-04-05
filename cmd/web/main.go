package main

import (
	"context"
	"finals-be/app/http"
	"finals-be/internal/config"
	"log"
)

func main() {

	cfg := config.Config{
		App: config.AppConfig{
			Name:     "finals-app",
			HttpPort: 8080,
			AppURL:   "http://localhost:8080",
		},
	}

	ctx := context.Background()

	server := http.NewServerOption(http.ServerOption{
		Config: &cfg,
	})

	err := server.Run(ctx, &cfg)
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
