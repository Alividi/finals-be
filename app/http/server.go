package http

import (
	"context"
	authService "finals-be/app/auth/service"
	userService "finals-be/app/user/service"
	"finals-be/internal/config"
	"finals-be/internal/connection"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

type Server struct {
	opts ServerOption

	authService *authService.AuthService
	userService *userService.UserService

	validate *validator.Validate
}

type ServerOption struct {
	Config *config.Config
	DB     *connection.SQLServerConnectionManager
}

func NewServerOption(opts ServerOption) Server {
	s := Server{
		opts: opts,
	}

	s.validate = validator.New()
	s.authService = authService.NewAuthService(opts.Config, opts.DB)
	s.userService = userService.NewUserService(opts.Config, opts.DB)

	return s
}

func (s *Server) Run(ctx context.Context, cfg *config.Config) error {

	router := RegisterRoutes(ctx, s, cfg)

	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONAL", "PATCH", "HEAD"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		MaxAge:             60,
		AllowCredentials:   true,
		OptionsPassthrough: false,
		Debug:              false,
	})

	httpHandler := c.Handler(router)

	if err := startServer(ctx, httpHandler, cfg); err != nil {
		return err
	}

	return nil
}

func startServer(ctx context.Context, httpHandler http.Handler, cfg *config.Config) error {
	errChan := make(chan error)

	go func() {
		errChan <- startHTTP(ctx, httpHandler, cfg)
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func startHTTP(ctx context.Context, httpHandler http.Handler, cfg *config.Config) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.HttpPort),
		Handler: httpHandler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info().Msg("Server is shutting down")
		}
	}()

	log.Info().Msgf("%s is starting at port: %d", cfg.App.Name, cfg.App.HttpPort)

	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-interruption

	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to shutdown")
		return err
	}

	log.Info().Msg("Server is shutdown")
	return nil
}
