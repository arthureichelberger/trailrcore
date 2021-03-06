//go:build unit
// +build unit

package store_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/arthureichelberger/trailrcore/user/model"
	"github.com/arthureichelberger/trailrcore/user/store"
	"github.com/stretchr/testify/assert"
)

func TestItShouldBeAbleToBuildTheFakeStore(t *testing.T) {
	fs := store.NewFakeStore()

	assert.IsType(t, *fs, store.FakeStore{})
	assert.Implements(t, new(store.Store), fs)
}

func TestFakeStoreCreateUserHandlerShouldBeConfigurable(t *testing.T) {
	fs := store.NewFakeStore()

	err := fs.CreateUser(context.Background(), model.User{})
	assert.NoError(t, err)

	fs.CreateUserHandler = func(ctx context.Context, user model.User) error {
		return fmt.Errorf("an error occurred")
	}
	err = fs.CreateUser(context.Background(), model.User{})
	assert.Error(t, err)
}

func TestFakeStoreGetUserByEmailHandlerShouldBeConfigurable(t *testing.T) {
	fs := store.NewFakeStore()

	user, err := fs.GetUserByEmail(context.Background(), "a@gmail.com")
	assert.Empty(t, user)
	assert.Error(t, err)

	fs.GetUserByEmailHandler = func(ctx context.Context, email string) (model.User, error) {
		return model.User{
			Email: "a@gmail.com",
		}, nil
	}

	user, err = fs.GetUserByEmailHandler(context.Background(), "a@gmail.com")
	assert.NotEmpty(t, user)
	assert.Equal(t, "a@gmail.com", user.Email)
	assert.NoError(t, err)
}
