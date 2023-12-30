package usecase

import (
	"context"
	"distributed_database_server/internal/auth/models"
)

type IUseCase interface {
	Register(ctx context.Context, params *models.SaveRequest) (*models.UserResponse, error)
	Login(ctx context.Context, params *models.LoginRequest) (*models.UserWithToken, error)
	// GetByEmail(ctx context.Context, email string) (*models.Response, error)
}
