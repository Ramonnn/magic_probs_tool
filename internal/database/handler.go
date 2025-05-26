package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type DatabaseHandler struct {
	Pool PgxPoolIface
}

// PgxPoolIface allows us to use both the real and mock pool
type PgxPoolIface interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

func NewDatabaseHandler() (*DatabaseHandler, error) {
	// Load environment variables from .env file (if exists)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or unable to load it")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PW")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	// Build connection string
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)

	// Setup connection pool config with timeout
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database config: %w", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.HealthCheckPeriod = time.Minute * 5

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	return &DatabaseHandler{
		Pool: pool,
	}, nil
}

// ExecuteQuery runs a query and returns rows (similar to Python execute_query)
func (db *DatabaseHandler) ExecuteQuery(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return rows, nil
}

// Close the connection pool when done
func (db *DatabaseHandler) Close() {
	if pool, ok := db.Pool.(*pgxpool.Pool); ok {
		pool.Close()
	}
}
