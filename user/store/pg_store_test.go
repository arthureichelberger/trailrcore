//go:build integration
// +build integration

package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/arthureichelberger/trailrcore/pkg/pgsql"
	"github.com/arthureichelberger/trailrcore/user/model"
	"github.com/arthureichelberger/trailrcore/user/store"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func getPgCnx() *sqlx.DB {
	return pgsql.Connect(context.Background(), "127.0.0.1", "5432", "trailrcore", "trailrcore", "trailrcore")
}

func teardown(db *sqlx.DB) {
	_ = db.MustExec(fmt.Sprintf("TRUNCATE %q;", "user"))
}

func TestItShouldBeAbleToBuildThePgStore(t *testing.T) {
	db := getPgCnx()
	defer teardown(db)
	pgStore := store.NewPgStore(db)

	assert.IsType(t, pgStore, store.PgStore{})
	assert.Implements(t, new(store.Store), pgStore)
}

func TestItShouldBeAbleToCreateAUser(t *testing.T) {
	db := getPgCnx()
	defer teardown(db)
	pgStore := store.NewPgStore(db)
	user := model.User{
		ID:        uuid.New(),
		Email:     "test@gmail.com",
		Password:  "test",
		CreatedAt: time.Now(),
	}

	err := pgStore.CreateUser(context.Background(), user)
	assert.NoError(t, err)
}

func TestItShouldNotBeAbleToCreateTwoUsersWithTheSameEmail(t *testing.T) {
	db := getPgCnx()
	defer teardown(db)
	pgStore := store.NewPgStore(db)
	user := model.User{
		ID:        uuid.New(),
		Email:     "test@gmail.com",
		Password:  "test",
		CreatedAt: time.Now(),
	}

	err := pgStore.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	err = pgStore.CreateUser(context.Background(), user)
	assert.Error(t, err)
}

func TestItShouldNotBeAbleToGetUserByEmailIfUserDoesNotExist(t *testing.T) {
	db := getPgCnx()
	defer teardown(db)
	pgStore := store.NewPgStore(db)

	user, err := pgStore.GetUserByEmail(context.Background(), "a@gmail.com")
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestItShouldBeAbleToGetUserByEmail(t *testing.T) {
	db := getPgCnx()
	defer teardown(db)
	pgStore := store.NewPgStore(db)
	userToInsert := model.User{
		ID:        uuid.New(),
		Email:     "a@gmail.com",
		Password:  "test",
		CreatedAt: time.Now(),
	}

	err := pgStore.CreateUser(context.Background(), userToInsert)
	assert.NoError(t, err)

	user, err := pgStore.GetUserByEmail(context.Background(), "a@gmail.com")
	assert.NotEmpty(t, userToInsert)
	assert.Equal(t, userToInsert.ID, user.ID)
	assert.NoError(t, err)
}
