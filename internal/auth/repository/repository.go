package repository

import (
	"database/sql"
	"fmt"

	"github.com/DragonPow/movie_booking/internal/auth/config"
	_ "github.com/lib/pq"
)

// Error checking helpers
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return err == sql.ErrNoRows
}

type Repository interface {
	Querier
}

type PostgresRepository struct {
	*Queries
	db *sql.DB
}

func NewPostgresRepository(cfg config.DatabaseConfig) (*PostgresRepository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &PostgresRepository{
		Queries: New(db),
		db:      db,
	}, nil
}

// GetDB returns the underlying database connection
func (r *PostgresRepository) GetDB() *sql.DB {
	return r.db
}

// Close closes the database connection
func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
