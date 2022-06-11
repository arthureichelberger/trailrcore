package pgsql

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres support.
	"github.com/rs/zerolog/log"
)

func Connect(ctx context.Context, host, port, user, pwd, db string) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, db)

	sqlxCnx, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Caller().Msg("Error while opening sqlx cnx")
	}

	sqlxCnx.SetConnMaxLifetime(time.Minute * 3)
	sqlxCnx.SetMaxOpenConns(10)
	sqlxCnx.SetMaxIdleConns(10)

	if err := sqlxCnx.PingContext(ctx); err != nil {
		log.Ctx(ctx).Panic().Err(err).Msg("Could not ping database")
	}

	return sqlxCnx
}
