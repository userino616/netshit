package config

import (
	"sync"


	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	DB struct {
		Host     string `env:"DB_HOST" env-required:"true"`
		Name     string `env:"DB_NAME" env-required:"true"`
		Port     string `env:"DB_PORT" env-default:"5432"`
		User     string `env:"DB_USER" env-required:"true"`
		Password string `env:"DB_PASSWORD" env-required:"true"`
	}

	Redis struct {
		Addr     string `env:"REDIS_ADDR" env-required:"true"`
		Password string `env:"REDIS_PASSWORD" env-required:"true"`
	}

	JWT struct {
		Secret                 string `env:"JWT_SECRET" env-required:"true"`
		AccessTokenExpiryHours uint   `env:"JWT_ACCESS_TOKEN_EXPIRY_HOURS" env-required:"true"`
	}

	Password struct {
		Secret string `env:"PASSWORD_SECRET" env-required:"true"`
	}

	Server struct {
		Addr        string `env:"SERVER_ADDR"  env-required:"true"`
		GRPCAddr    string `env:"GRPC_ADDR" env-required:"true"`
		GRPCTimeout uint   `env:"GRPC_TIMEOUT" env-required:"true"`
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
