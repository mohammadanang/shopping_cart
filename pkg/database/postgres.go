package database

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, conf *config.Config) (*dbgen.Queries, *pgxpool.Pool) {
	dbConnStr := "postgres://" + conf.DBUser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + strconv.Itoa(conf.DBPort) + "/" + conf.DBName + "?sslmode=disable"
	cfg, err := pgxpool.ParseConfig(dbConnStr)
	if err != nil {
		log.Fatalf("failed to parse dsn: %v", err)
	}

	// optional tuning
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return dbgen.New(pool), pool
}
