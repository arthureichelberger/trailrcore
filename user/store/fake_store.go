package store

import (
	"context"

	"github.com/arthureichelberger/trailrcore/user/model"
)

type FakeStore struct {
	CreateUserHandler func(ctx context.Context, user model.User) error
}

func NewFakeStore() FakeStore {
	return FakeStore{
		CreateUserHandler: func(ctx context.Context, user model.User) error {
			return nil
		},
	}
}

func (fs FakeStore) CreateUser(ctx context.Context, user model.User) error {
	return fs.CreateUserHandler(ctx, user)
}
