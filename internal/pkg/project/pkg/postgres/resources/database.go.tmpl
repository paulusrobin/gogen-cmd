package postgres

import (
	"github.com/rs/zerolog/log"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDatabase initialize postgres database.
func NewDatabase(database Config) (*gorm.DB, error) {
	log.Debug().Msg("Connecting to Postgres")

	// Init gormDB connection
	db, err := gorm.Open(pg.Open(database.DSN), &gorm.Config{
		Logger:                 logger.Default.LogMode(database.LogLevel),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Error().Err(err).Msg("could not connect to db")
		return nil, err
	}

	// Configure conn pool through the underlying *sql.DB instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("could not configure db conn pool")
		return nil, err
	}

	// Sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(database.MaxOpenConnections)
	// Sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(database.MaxIdleConnections)
	// Sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(database.MaxConnectionLifetime)

	return db, nil
}
