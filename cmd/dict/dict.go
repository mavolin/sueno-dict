package main

import (
	"go.uber.org/zap"

	"github.com/mavolin/sueno-dict/internal/config"
	"github.com/mavolin/sueno-dict/internal/setup/logger"
	"github.com/mavolin/sueno-dict/repository/postgres"
)

func main() {
	log, err := logger.Setup()
	if err != nil {
		panic(err)
	}

	err = run(log)
	if err != nil {
		log.Fatal(err)
	}
}

func run(log *zap.SugaredLogger) error {
	c, err := config.Read()
	if err != nil {
		return err
	}

	repo, err := postgres.Open(postgres.Options{
		DBName:   c.DB.Name,
		Host:     c.DB.Host,
		User:     c.DB.User,
		Password: c.DB.Password,
	})
	if err != nil {
		return err
	}

	return nil
}
