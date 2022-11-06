package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type Config struct {
	Server struct {
		Port string `env:"PORT" envDefault:"8080"`
	} `envPrefix:"SERVER_"`

	DB struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		Name     string `env:"NAME" envDefault:"pengpeng"`
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
