package repository

import (
	"context"
	"distributed_database_server/internal/auth/entity"
)

// Auth repository interface
type IRepository interface {
	// Register(ctx context.Context, user *entity.User) (int64, error)
	Create(ctx context.Context, todo *entity.User) (*entity.User, error)
	// Update(ctx context.Context, user *models.User) (*models.User, error)
	// Delete(ctx context.Context, userID uuid.UUID) error
	GetById(ctx context.Context, userId int) (*entity.User, error)
	// FindByEmail(ctx context.Context, email string) (*entity.User, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.User, error)
}

// Auth Redis repository interface
type IRedisRepository interface {
	GetById(ctx context.Context, key string) (*entity.User, error)
	SetUser(ctx context.Context, key string, seconds int, user *entity.User) error
	// DeleteUser(ctx context.Context, key string) error
}
