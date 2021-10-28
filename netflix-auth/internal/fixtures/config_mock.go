package fixtures

import "netflix-auth/internal/config"

var CFG = &config.Config{
	DB: struct {
		Host     string `env:"DB_HOST" env-required:"true"`
		Name     string `env:"DB_NAME" env-required:"true"`
		Port     string `env:"DB_PORT" env-required:"true"`
		User     string `env:"DB_USER" env-required:"true"`
		Password string `env:"DB_PASSWORD" env-required:"true"`
	}{
		Host:     "localhost",
		Name:     "netflix_test",
		Port:     "5432",
		User:     "postgres",
		Password: "qwerty",
	},
}
