package db

import (
	"alter-io-go/config"
	"alter-io-go/helpers/logger"
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabaseConnection(cfg *config.AppConfig) *pgxpool.Pool {
	var (
		defaultMaxConns          int32         = cfg.Database.MaxConns
		defaultMinConns          int32         = cfg.Database.MinConns
		defaultMaxConnLifetime   time.Duration = time.Duration(cfg.Database.MaxConnLifeTime) * time.Minute
		defaultMaxConnIdleTime   time.Duration = time.Duration(cfg.Database.MaxConnIdleTime) * time.Minute
		defaultHealthCheckPeriod time.Duration = time.Duration(cfg.Database.HealthCheckPeriod) * time.Second
		defaultConnTimeout       time.Duration = time.Duration(cfg.Database.ConnTimeout) * time.Second
	)

	// Load database configuration from environment variables
	// example: postgres://user:password@localhost:5432/dbname?sslmode=disable
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Address,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	dbConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		logger.Get().With().ErrorContext(context.Background(), "Failed to create database config", slog.Any("error", err))
		os.Exit(1)
	}

	// Apply pooling settings
	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnTimeout

	// Create the connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		logger.Get().With().ErrorContext(context.Background(), "Failed to establish database connection", slog.Any("error", err))
		os.Exit(1)
	}

	logger.Get().With().Info("Connected to the database successfully!")
	return pool
}
