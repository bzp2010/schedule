package database

import (
	"path/filepath"
	"strings"

	"github.com/bzp2010/schedule/internal/config"
	"github.com/bzp2010/schedule/internal/database/models"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	dsnPrefixSqlite = "sqlite://"
)

var (
	database *gorm.DB
)

// GetDatabase gets the `gorm.DB` singleton
func GetDatabase() *gorm.DB {
	return database
}

// SetupDatabase will create a single instance of `gorm.DB` based on the configuration
func SetupDatabase(cfg config.Config) error {
	var (
		dsn = cfg.DSN
		err error
	)

	// sqlite
	if strings.HasPrefix(dsn, dsnPrefixSqlite) {
		path := strings.TrimPrefix(dsn, dsnPrefixSqlite)
		path, err := filepath.Abs(path)
		if err != nil {
			return errors.Wrap(err, "failed to get absolute path to sqlite database file")
		}

		database, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	}

	// no instance of any kind of database is created
	if database == nil {
		return errors.New("unsupport database type")
	}

	if err != nil {
		return errors.Wrap(err, "failed to open sqlite database")
	}

	if err := autoMigrate(database); err != nil {
		return errors.Wrap(err, "failed to auto migrate database model")
	}

	return nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(models.Task{})
}
