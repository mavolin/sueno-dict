package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup() (*zap.SugaredLogger, error) {
	log, err := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.WarnLevel), zap.AddCaller())
	if err != nil {
		return nil, errors.Wrap(err, "setup: RegularLogger")
	}

	return log.Sugar(), nil
}
