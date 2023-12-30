package usecase

import (
	"context"
	"distributed_database_server/config"
	"distributed_database_server/internal/auth/entity"
	"distributed_database_server/internal/auth/models"
	"distributed_database_server/internal/auth/repository"
	"distributed_database_server/internal/constants"
	"distributed_database_server/package/utils"
	"fmt"

	"github.com/labstack/gommon/log"
)

type usecase struct {
	cfg       *config.Config
	repo      repository.IRepository
	redisRepo repository.IRedisRepository
}

// Constructor
func NewUseCase(cfg *config.Config, repo repository.IRepository, redisRepo repository.IRedisRepository) IUseCase {
	return &usecase{
		cfg:       cfg,
		repo:      repo,
		redisRepo: redisRepo,
	}
}

const (
	basePrefix = "api-auth:"
)

func (u *usecase) GenerateUserKey(userId int) string {
	return fmt.Sprintf("%s: %d", basePrefix, userId)
}

func (u *usecase) Register(ctx context.Context, params *models.SaveRequest) (*models.UserResponse, error) {
	log.SetPrefix("[Register]")
	log.Infof("Register user with params: {FirstName: %s, LastName: %s, UserName: %s}",
		params.FirstName, params.LastName, params.UserName)

	// validation

	// check if user already exists
	foundUser, err := u.repo.GetOne(ctx, (&models.RequestList{UserName: params.UserName}).ToMap())
	if err != nil {
		log.Errorf("Error while finding user by email: %s", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser.Id != 0 {
		log.Errorf("User already exist with email: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_EMAIL_ALREADY_EXISTS)
	}

	if params.Gender != "Male" && params.Gender != "Female" {
		log.Errorf("Invalid gender type: %s", params.Gender)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INVALID_GENDER_TYPE)
	}
	// end validation

	// create new user
	obj := &entity.User{}
	obj.HashPassword()
	obj.ParseFromSaveRequest(params)
	res, err := u.repo.Create(ctx, obj)
	if err != nil {
		log.Errorf("Error while creating new user: %s", err)
		return nil, err
	}
	res.SanitizePassword()
	return res.Export(), nil
}

func (u *usecase) Login(ctx context.Context, params *models.LoginRequest) (*models.UserWithToken, error) {
	log.SetPrefix("[Login]")
	log.Infof("Sign in with user {UserName: %v}", params.UserName)

	// validation

	// check if user already exists
	foundUser, err := u.repo.GetOne(ctx, (&models.RequestList{UserName: params.UserName}).ToMap())
	if err != nil {
		log.Errorf("Error while finding user by email: %s", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser == nil {
		log.Errorf("User not found with email: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_USER_NOT_FOUND)
	}

	// check if password is correct
	if err = utils.ComparePasswords(foundUser.Password, params.Password); err != nil {
		log.Errorf("Compare password failed: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_UNAUTHORIZED, constants.STATUS_MESSAGE_INVALID_EMAIL_OR_PASSWORD)
	}
	// end validation

	// generate token
	token, err := utils.GenerateJWTToken(foundUser.Export(), u.cfg.Auth.JWTSecret, u.cfg.Auth.Expire)
	if err != nil {
		log.Errorf("Cannot generate token: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	// save to cache
	if err = u.redisRepo.SetUser(ctx, u.GenerateUserKey(foundUser.Id), u.cfg.Auth.Expire, foundUser); err != nil {
		log.Errorf("usecase.redisRepo.SetUser: %v", err)
		return nil, err
	}

	foundUser.SanitizePassword()

	return &models.UserWithToken{
		User:  foundUser.Export(),
		Token: token,
	}, nil
}
