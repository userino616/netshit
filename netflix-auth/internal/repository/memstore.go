package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const BlackListTag = "blackList"

type MemStore interface {
	BlackListJWT(tokenID string, exp time.Duration) error
	IsJWTBlacklisted(tokenID string) (bool, error)
}

type redisRepository struct {
	*redis.Client
}

func NewMemoryStorage(client *redis.Client) MemStore {
	return redisRepository{client}
}

func (r redisRepository) BlackListJWT(tokenID string, exp time.Duration) error {
	key :=  fmt.Sprintf("%s:%s", BlackListTag, tokenID)
	err := r.Set(context.Background(), key, 0, exp).Err()

	return err
}


func (r redisRepository) IsJWTBlacklisted(tokenID string) (bool, error) {
	key :=  fmt.Sprintf("%s:%s", BlackListTag, tokenID)
	_, err := r.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}

	return true, err
}
