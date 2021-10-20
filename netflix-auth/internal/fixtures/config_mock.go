package fixtures

import "netflix-auth/internal/config"

var CFG = &config.Config{
	DB: struct {
		Host     string `yaml:"host" env-required:"true"`
		Name     string `yaml:"name" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		User     string `yaml:"user" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
	}{
		Host:     "localhost",
		Name:     "netflix_test",
		Port:     "5432",
		User:     "postgres",
		Password: "qwerty",
	},
}
