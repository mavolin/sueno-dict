package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type Config struct {
	Server struct {
		Port string `env:"PORT" envDefault:"8080"`

		ProtectedUser     string `env:"PROTECTED_USER" envRequired:"true"`
		ProtectedPassword string `env:"PROTECTED_PASSWORD" envRequired:"true"`
	} `envPrefix:"SERVER_"`

	DB struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		Name     string `env:"NAME" envDefault:"sueno_dict"`
		User     string `env:"USER" envDefault:"root"`
		Password string `env:"PASSWD" envDefault:"root"`
	} `envPrefix:"DB_"`
}

func Read() (*Config, error) {
	var c Config

	err := env.Parse(&c, env.Options{Prefix: "SUENO_DICT_"})
	if err != nil {
		return nil, errors.Wrap(err, "could not read config")
	}

	return &c, nil
}
