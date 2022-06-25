package store

import (
	"context"
	"fmt"

	"github.com/arthureichelberger/trailrcore/user/model"
)

type FakeStore struct {
	CreateUserHandler     func(ctx context.Context, user model.User) error
	GetUserByEmailHandler func(ctx context.Context, email string) (model.User, error)
}

func NewFakeStore() *FakeStore {
	return &FakeStore{
		CreateUserHandler: func(ctx context.Context, user model.User) error {
			return nil
		},
		GetUserByEmailHandler: func(ctx context.Context, email string) (model.User, error) {
			return model.User{}, fmt.Errorf("no user found")
		},
	}
}

func (fs *FakeStore) CreateUser(ctx context.Context, user model.User) error {
	return fs.CreateUserHandler(ctx, user)
}

func (fs *FakeStore) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	return fs.GetUserByEmailHandler(ctx, email)
}
