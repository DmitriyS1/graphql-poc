package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewPostgresDb(ctx context.Context, dbUrl string) DB {
	dbConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalln("Unable to parse database URL:", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	return DB{Pool: pool}
}
