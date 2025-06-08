package connection

import (
	"finals-be/internal/config"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
)

type SQLServerConnectionManager struct {
	db *sqlx.DB
	m  *migrate.Migrate
}

func NewConnectionManager(cfg config.DatabaseConfig) (*SQLServerConnectionManager, error) {
	db, err := sqlx.Connect(cfg.Driver, cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetConnMaxIdleTime(cfg.MaxIdleDuration)
	db.SetConnMaxLifetime(cfg.MaxLifeTimeDuration)

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create mysql driver instance: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		cfg.MigrationPath,
		cfg.Driver,
		driver,
	)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create migration instance: %w", err)
	}

	log.Info().Msg("database connection established")

	return &SQLServerConnectionManager{
		db: db,
		m:  m,
	}, nil
}

func (cm *SQLServerConnectionManager) RunMigration() error {
	err := cm.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	log.Info().Msg("database migration completed")
	return nil
}

func (cm *SQLServerConnectionManager) Close() error {
	log.Info().Msg("closing database connection")
	return cm.db.Close()
}

func (cm *SQLServerConnectionManager) GetQuery() *SingleInstruction {
	return NewSingleInstruction(cm.db)
}

func (cm *SQLServerConnectionManager) GetTransaction() *MultiInstruction {
	return NewMultiInstruction(cm.db)
}
