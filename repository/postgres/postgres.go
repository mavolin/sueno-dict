// Package postgres provides a postgresql implementation of the
// [repository.Repository].
package postgres

import (
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"

	"github.com/mavolin/sueno-dict/repository"
)

var _ repository.Repository = (*Repository)(nil)

type Repository struct {
	DB *gorm.DB
}

// Options contains the options used to configure the repository.
type Options struct {
	// DBName is the name of the database to use.
	//
	// It must exist.
	//
	// It is required.
	DBName string

	// Host is the host of the database.
	//
	// It is required.
	Host string

	// User is the username to use.
	//
	// It is required.
	User string
	// Password is the password to use.
	//
	// It is required.
	Password string

	// Logger is the logger used to log sql queries.
	Logger *zap.SugaredLogger
}

func Open(o Options) (*Repository, error) {
	var dsnBuilder strings.Builder

	dsnBuilder.WriteString("postgres://")
	dsnBuilder.WriteString(o.User)
	dsnBuilder.WriteString(":")
	dsnBuilder.WriteString(o.Password)
	dsnBuilder.WriteString("@")
	dsnBuilder.WriteString(o.Host)
	dsnBuilder.WriteString("/")
	dsnBuilder.WriteString(o.DBName)

	dsn := dsnBuilder.String()
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, errors.Wrap(err, "postgres: could not open database")
	}

	log := zapgorm2.New(o.Logger.Named("gorm").Desugar())
	log.LogLevel = logger.Info
	log.IgnoreRecordNotFoundError = true
	db.Logger = log

	repo := &Repository{DB: db}

	if err = repo.migrate(); err != nil {
		return nil, err
	}

	return repo, nil
}
