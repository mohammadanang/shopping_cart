package database

import (
	"context"
	"strconv"
	"time"

	"github.com/mohammadanang/shopping-cart/db/dbgen"
	"github.com/mohammadanang/shopping-cart/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Queries *dbgen.Queries
	Pool    *pgxpool.Pool
	Err     error
}

func NewPostgres(ctx context.Context, conf *config.Config) <-chan Postgres {
	result := make(chan Postgres)

	go func() {
		defer close(result)

		dbConnStr := "postgres://" + conf.DBUser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + strconv.Itoa(conf.DBPort) + "/" + conf.DBName + "?sslmode=" + conf.SSLMode
		cfg, err := pgxpool.ParseConfig(dbConnStr)
		if err != nil {
			result <- Postgres{Err: err}
			return
		}

		// optional tuning
		cfg.MaxConns = 10
		cfg.MinConns = 2
		cfg.MaxConnLifetime = time.Hour

		pool, err := pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			result <- Postgres{Err: err}
			return
		}

		result <- Postgres{
			Queries: dbgen.New(pool),
			Pool:    pool,
		}
	}()

	return result
}
