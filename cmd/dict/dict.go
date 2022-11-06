package main

import (
	"go.uber.org/zap"

	"github.com/mavolin/sueno-dict/internal/setup/logger"
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
	return nil
}
