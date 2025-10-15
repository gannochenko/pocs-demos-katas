package database

import (
	"gateway/internal/domain"
	"log/slog"
	"sync"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type Database struct {
	DB        *gorm.DB
	config    *domain.DatabaseConfig
	logger    *slog.Logger
	closeOnce sync.Once
}

func NewDatabase(config *domain.DatabaseConfig, logger *slog.Logger) *Database {
	return &Database{
		config: config,
		logger: logger,
	}
}

func (d *Database) Connect() (func() error, error) {
	db, err := gorm.Open(postgres.Open(d.config.DSN), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	d.DB = db
	d.logger.Info("connected to database")

	return func() error {
		var closeErr error
		d.closeOnce.Do(func() {
			sqlDB, err := db.DB()
			if err != nil {
				closeErr = errors.Wrap(err, "failed to get database instance")
				return
			}

			if err := sqlDB.Close(); err != nil {
				closeErr = errors.Wrap(err, "failed to close database connection")
				return
			}

			d.logger.Info("database connection closed")
		})
		return closeErr
	}, nil
}
