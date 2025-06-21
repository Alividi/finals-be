package service

import (
	"context"
	"finals-be/app/auth/dto"
	userRepo "finals-be/app/user/repository"
	"finals-be/internal/config"
	"finals-be/internal/connection"
	"finals-be/internal/lib/auth"
	"finals-be/internal/lib/helper"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	cfg            *config.Config
	db             *connection.MultiInstruction
	userRepository userRepo.IUserRepository
}

func NewAuthService(cfg *config.Config, conn *connection.SQLServerConnectionManager) *AuthService {
	db := conn.GetTransaction()
	return &AuthService{
		cfg:            cfg,
		db:             db,
		userRepository: userRepo.NewUserRepository(db),
	}
}

func (a *AuthService) Login(ctx context.Context, request dto.LoginRequest) (response dto.LoginResponse, err error) {

	user, err := a.userRepository.GetByUsername(ctx, request.Username)
	if err != nil {
		return
	}

	if helper.ComparePassword(user.Password, request.Password) {
		return response, helper.NewErrUnauthorized("Username or Password invalid")
	}

	accessTokenClaims := auth.NewAuthClaims(user.ID, user.Username, user.Role, a.cfg.App.Name, time.Now().Add(a.cfg.JWT.LoginExpirationDuration))
	accessToken, err := auth.GenerateToken(accessTokenClaims, &a.cfg.JWT)
	if err != nil {
		log.Default().Println("Failed to generate accessToken")
		return
	}

	refreshTokenClaims := auth.NewAuthClaims(user.ID, user.Username, user.Role, a.cfg.App.Name, time.Now().Add(a.cfg.JWT.RefreshExpirationDuration))
	refreshToken, err := auth.GenerateToken(refreshTokenClaims, &a.cfg.JWT)
	if err != nil {
		log.Default().Println("Failed to generate refreshToken")
		return
	}

	if err = a.userRepository.StoreRefreshToken(ctx, user.Username, &refreshToken); err != nil {
		log.Default().Println("Failed to store refresh token")
		return
	}

	response = dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, err
}

func (a *AuthService) RefreshToken(ctx context.Context, request dto.RefreshTokenRequest) (response dto.LoginResponse, err error) {

	user, err := a.userRepository.GetByRefreshToken(ctx, request.RefreshToken)
	if err != nil {
		log.Default().Println("Failed to find user")
		return
	}

	refreshTokenClaims, err := auth.ValidateToken(&a.cfg.JWT, *user.RefreshToken)
	if err != nil {
		log.Default().Println("Invalid token")
		return
	}

	refreshTokenClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(a.cfg.JWT.LoginExpirationDuration))

	accessTokenClaims := auth.NewAuthClaims(user.ID, user.Username, user.Role, a.cfg.App.Name, time.Now().Add(a.cfg.JWT.LoginExpirationDuration))
	accessToken, err := auth.GenerateToken(accessTokenClaims, &a.cfg.JWT)
	if err != nil {
		log.Default().Println("Failed to generate accessToken")
		return
	}

	refreshToken, err := auth.GenerateToken(refreshTokenClaims, &a.cfg.JWT)
	if err != nil {
		log.Default().Println("Failed to generate refreshToken")
		return
	}

	if err = a.userRepository.StoreRefreshToken(ctx, user.Username, &refreshToken); err != nil {
		log.Default().Println("Failed to store refresh token")
		return
	}

	response = dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, err
}

func (a *AuthService) Logout(ctx context.Context, request dto.LogoutRequest) error {

	user := auth.GetUserContext(ctx)

	a.userRepository.StoreRefreshToken(ctx, user.Username, nil)

	// TODO: implement remove FcmToken

	return nil
}
