package repository

import (
	"context"
	"distributed_database_server/internal/auth/entity"
	"encoding/json"
	"time"
)

// Get user by id
func (r *redisRepo) GetById(ctx context.Context, key string) (*entity.User, error) {

	userBytes, err := r.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	user := &entity.User{}
	if err = json.Unmarshal(userBytes, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Cache user with duration in seconds
func (r *redisRepo) SetUser(ctx context.Context, key string, seconds int, user *entity.User) error {

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err = r.redisClient.Set(ctx, key, userBytes, (time.Second * time.Duration(seconds))).Err(); err != nil {
		return err
	}
	return nil
}
