package postgres

import (
	"context"
	"netflix-movies/internal/config"
	"sync"

	"github.com/go-pg/pg/v10"
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
		//dsn, err := pg.ParseURL("postgresql://postgres:qwerty@localhost/netflix?sslmode=disable")
		//if err != nil {
		//	panic(err)
		//}
		db = pg.Connect(dbOptions)
		if err := db.Ping(context.Background()); err != nil {
			panic(err)
		}
	})
}

func GetDB() *pg.DB {
	return db
}
