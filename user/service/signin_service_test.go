//go:build unit
// +build unit

package service_test

import (
	"context"
	"testing"

	"github.com/arthureichelberger/trailrcore/user/model"
	"github.com/arthureichelberger/trailrcore/user/service"
	"github.com/arthureichelberger/trailrcore/user/store"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestItShouldNotBeAbleToSigninWithAnUnknownUser(t *testing.T) {
	fs := store.NewFakeStore()
	sis := service.NewSigninService(fs)

	user, err := sis.Login(context.Background(), "a@gmail.com", "test")
	assert.Error(t, err)
	assert.Empty(t, user)
}

func TestItShouldNotBeAbleToSigninWithTheWrongPassword(t *testing.T) {
	fs := store.NewFakeStore()
	pwd, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	fs.GetUserByEmailHandler = func(ctx context.Context, email string) (model.User, error) {
		return model.User{
			Password: string(pwd),
			Email:    "a@gmail.com",
		}, nil
	}
	sis := service.NewSigninService(fs)

	user, err := sis.Login(context.Background(), "a@gmail.com", "testtest")
	assert.Error(t, err)
	assert.Empty(t, user)
}

func TestItShouldBeAbleToSignin(t *testing.T) {
	fs := store.NewFakeStore()
	pwd, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	fs.GetUserByEmailHandler = func(ctx context.Context, email string) (model.User, error) {
		return model.User{
			Password: string(pwd),
			Email:    "a@gmail.com",
		}, nil
	}
	sis := service.NewSigninService(fs)

	user, err := sis.Login(context.Background(), "a@gmail.com", "test")
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
