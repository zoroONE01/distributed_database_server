package init

import (
	"distributed_database_server/config"
	handler "distributed_database_server/internal/auth/delivery/http"
	"distributed_database_server/internal/auth/repository"
	"distributed_database_server/internal/auth/usecase"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Init struct {
	Repository repository.IRepository
	Usecase    usecase.IUseCase
	Handler    handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	redisClient *redis.Client,
) *Init {
	repo := repository.NewRepo(db)
	redisRepo := repository.NewRedisRepo(redisClient)
	usecase := usecase.NewUseCase(cfg, repo, redisRepo)
	handler := handler.NewHandler(cfg, usecase)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
