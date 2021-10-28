package config

import (
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	DB struct {
		Host     string `env:"DB_HOST" env-required:"true"`
		Name     string `env:"DB_NAME" env-required:"true"`
		Port     string `env:"DB_PORT" env-required:"true"`
		User     string `env:"DB_USER" env-required:"true"`
		Password string `env:"DB_PASSWORD" env-required:"true"`
	}

	Server struct {
		Addr string `env:"SERVER_ADDR"  env-required:"true"`
	}

	LogLvl logrus.Level `env:"LOG_LEVEL" env-required:"true"`
}

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadEnv(instance)
		if err != nil {
			panic(err)
		}
	})

	return instance
}
