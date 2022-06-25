package store

import (
	"context"
	"fmt"

	"github.com/arthureichelberger/trailrcore/user/model"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type PgStore struct {
	db *sqlx.DB
}

func NewPgStore(db *sqlx.DB) PgStore {
	return PgStore{
		db: db,
	}
}

func (ps PgStore) CreateUser(ctx context.Context, user model.User) error {
	query := `INSERT INTO "user" (id, email, password, created_at) VALUES (:id, :email, :password, :created_at);`

	rows, err := ps.db.NamedQueryContext(ctx, query, user)
	if err != nil {
		log.Error().Err(err).Str("email", user.Email).Msg("could not insert user")
		return fmt.Errorf("could not insert user: %w", err)
	}

	defer rows.Close()

	return nil
}

func (ps PgStore) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	query := `SELECT id, email, password FROM "user" WHERE email = $1;`

	var user model.User
	if err := ps.db.GetContext(ctx, &user, query, email); err != nil {
		log.Error().Err(err).Str("email", email).Msg("could not get user")
		return model.User{}, fmt.Errorf("could not get user: %w", err)
	}

	return user, nil
}
