//go:build unit
// +build unit

package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/arthureichelberger/trailrcore/user/exception"
	"github.com/arthureichelberger/trailrcore/user/model"
	"github.com/arthureichelberger/trailrcore/user/service"
	"github.com/arthureichelberger/trailrcore/user/store"
	"github.com/stretchr/testify/assert"
)

func TestItShouldBeAbleToBuildTheSignInService(t *testing.T) {
	s := service.NewSignInService(store.NewFakeStore())
	assert.IsType(t, s, service.SignInService{})
}

func TestItShouldNotBeAbleToCreateAUserWithAnInvalidEmail(t *testing.T) {
	fs := store.NewFakeStore()
	s := service.NewSignInService(fs)

	user, err := s.CreateUser(context.Background(), "", "", "")
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.ErrorAs(t, err, new(exception.CouldNotValidateEmailError))
}

func TestItShouldNotBeAbleToCreateAUserWithAnEmptyPassword(t *testing.T) {
	fs := store.NewFakeStore()
	s := service.NewSignInService(fs)

	user, err := s.CreateUser(context.Background(), "test@gmail.com", "", "")
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.ErrorAs(t, err, new(exception.CouldNotValidatePasswordError))
}

func TestItShouldNotBeAbleToCreateAUserWithAnInvalidPassword(t *testing.T) {
	fs := store.NewFakeStore()
	s := service.NewSignInService(fs)

	user, err := s.CreateUser(context.Background(), "test@gmail.com", "a", "")
	assert.Empty(t, user)
	assert.Error(t, err)
	assert.ErrorAs(t, err, new(exception.CouldNotValidatePasswordError))
}

func TestItShouldFailGracefullyIfItIsNotPossibleToInsertUserInDatabase(t *testing.T) {
	fs := store.NewFakeStore()
	fs.CreateUserHandler = func(ctx context.Context, user model.User) error {
		return fmt.Errorf("an error occurred")
	}
	s := service.NewSignInService(fs)

	user, err := s.CreateUser(context.Background(), "test@gmail.com", "a", "a")
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestItShouldBeAbleToCreateAUser(t *testing.T) {
	fs := store.NewFakeStore()
	s := service.NewSignInService(fs)

	user, err := s.CreateUser(context.Background(), "test@gmail.com", "a", "a")
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
}
