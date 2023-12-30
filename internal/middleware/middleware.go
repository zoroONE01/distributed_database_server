package middleware

import (
	"distributed_database_server/config"
	authRepo "distributed_database_server/internal/auth/repository"
)

// Middleware manager
type MiddlewareManager struct {
	cfg      *config.Config
	authRepo authRepo.IRepository
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, authRepo authRepo.IRepository) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:      cfg,
		authRepo: authRepo,
	}
}
