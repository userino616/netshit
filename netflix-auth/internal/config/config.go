package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	DB struct {
		Host     string `yaml:"host" env-required:"true"`
		Name     string `yaml:"name" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		User     string `yaml:"user" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
	} `yaml:"db" env-required:"true"`

	JWT struct {
		Secret                 string `yaml:"secret" env-required:"true"`
		AccessTokenExpiryHours uint   `yaml:"accessTokenExpiryHours" env-required:"true"`
	} `yaml:"jwt" env-required:"true"`

	Password struct {
		Secret string `yaml:"secret" env-required:"true"`
	} `yaml:"password" env-required:"true"`

	Server struct {
		Addr        string `yaml:"addr"  env-required:"true"`
		GRPCAddr    string `yaml:"grpcAddr" env-required:"true"`
		GRPCTimeout uint   `yaml:"grpcTimeout" env-required:"true"`
	} `yaml:"server" env-required:"true"`
}

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadConfig("config.yml", instance)
		if err != nil {
			panic(err)
		}
	})

	return instance
}
