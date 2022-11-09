package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/mavolin/sueno-dict/internal/meta"
)

func Setup() (*zap.SugaredLogger, error) {
	log, err := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.WarnLevel), zap.AddCaller())
	if err != nil {
		return nil, errors.Wrap(err, "setup: RegularLogger")
	}

	if !meta.Development {
		log = log.WithOptions(zap.IncreaseLevel(zapcore.InfoLevel))
	}

	return log.Sugar(), nil
}
