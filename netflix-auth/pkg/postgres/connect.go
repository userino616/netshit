package postgres

import (
	"context"
	"sync"

	"github.com/go-pg/pg/v10"
	"netflix-auth/internal/config"
)

var (
	db   *pg.DB
	once sync.Once
)

func Load(cfg *config.Config) {
	once.Do(func() {
		dbOptions := &pg.Options{
			Addr:     cfg.DB.Host + ":" + cfg.DB.Port,
			User:     cfg.DB.User,
			Password: cfg.DB.Password,
			Database: cfg.DB.Name,
		}
		db = pg.Connect(dbOptions)
		if err := db.Ping(context.Background()); err != nil {
			panic(err)
		}
	})
}

func GetDB() *pg.DB {
	return db
}

func Close() error {
	return db.Close()
}
