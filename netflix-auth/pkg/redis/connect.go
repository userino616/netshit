package redis

import (
	"github.com/go-redis/redis/v8"
	"netflix-auth/internal/config"
	"sync"
)

var (
	db   *redis.Client
	once sync.Once
)

func Load(cfg *config.Config) {
	once.Do(func() {
		dbOptions := &redis.Options{
			Addr:               cfg.Redis.Addr,
			Password:           cfg.Redis.Password,
			DB:                 0,
		}
		db = redis.NewClient(dbOptions)
	})
}

func GetDB() *redis.Client {
	return db
}

func Close() error {
	if conn := GetDB(); conn != nil {
		return db.Close()
	}
	return nil
}