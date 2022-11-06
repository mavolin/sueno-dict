package main

import (
	"os"
	"os/signal"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/mavolin/sueno-dict/internal/config"
	"github.com/mavolin/sueno-dict/internal/setup/logger"
	"github.com/mavolin/sueno-dict/internal/setup/router"
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
	setupLog := log.Named("setup")

	setupLog.Info("reading config")

	c, err := config.Read()
	if err != nil {
		return err
	}

	setupLog.
		With("config", c).
		Info("read config")

	setupLog.Info("connecting to database")

	repo, err := postgres.Open(postgres.Options{
		DBName:   c.DB.Name,
		Host:     c.DB.Host,
		User:     c.DB.User,
		Password: c.DB.Password,
	})
	if err != nil {
		return err
	}

	setupLog.Info("setting up router")

	r, err := router.Setup(router.Options{
		Repository:        repo,
		ProtectedUser:     c.Server.ProtectedUser,
		ProtectedPassword: c.Server.ProtectedPassword,
	})
	if err != nil {
		return err
	}

	if err = r.Run(":" + c.Server.Port); err != nil {
		return errors.Wrap(err, "start server")
	}

	setupLog.Info("👂 listening on port " + c.Server.Port)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	setupLog.Info("received SIGINT, shutting down")
	setupLog.Info("👋 bye")

	return nil
}
