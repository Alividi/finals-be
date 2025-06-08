package service

import (
	"context"
	"finals-be/app/user/dto"
	userRepo "finals-be/app/user/repository"
	"finals-be/internal/config"
	"finals-be/internal/connection"
)

type UserService struct {
	cfg            *config.Config
	db             *connection.MultiInstruction
	userRepository userRepo.IUserRepository
}

func NewUserService(cfg *config.Config, conn *connection.SQLServerConnectionManager) *UserService {
	db := conn.GetTransaction()
	return &UserService{
		cfg:            cfg,
		db:             db,
		userRepository: userRepo.NewUserRepository(db),
	}
}

func (u *UserService) GetCurrentUser(ctx context.Context, userId int64) (response dto.GetCurrentUserResponse, err error) {

	user, err := u.userRepository.GetUserDetail(ctx, userId)
	almt, _ := u.userRepository.GetAlamatUser(ctx, userId)

	if err != nil {
		return
	}

	response = dto.GetCurrentUserResponse{
		Role:     user.Role,
		Username: user.Username,
		Phone:    user.NoTelp,
		Email:    user.Email,
		Address:  almt,
	}

	return response, nil
}

func (u *UserService) GetTechnicians(ctx context.Context) (response []dto.GetTechniciansResponse, err error) {
	technicians, err := u.userRepository.GetTechnicians(ctx)
	if err != nil {
		return
	}

	for _, tech := range technicians {
		response = append(response, dto.GetTechniciansResponse{
			ID:     tech.ID,
			Nama:   tech.Nama,
			Email:  tech.Email,
			NoTelp: tech.NoTelp,
			Status: tech.Status,
			Base:   tech.Base,
		})
	}

	return response, nil
}

func (u *UserService) GetUserStatus(ctx context.Context, userId int64) (response dto.UserStatus, err error) {
	status, err := u.userRepository.GetUserStatus(ctx, userId)
	if err != nil {
		return
	}

	response = dto.UserStatus{
		NotificationCount: status.NotificationCount,
	}

	return response, nil
}
