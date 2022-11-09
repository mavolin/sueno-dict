// Package postgres provides a postgresql implementation of the
// [repository.Repository].
package postgres

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

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

	db.Logger = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	repo := &Repository{DB: db}

	if err = repo.migrate(); err != nil {
		return nil, err
	}

	return repo, nil
}
