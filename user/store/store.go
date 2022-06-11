package store

import (
	"context"

	"github.com/arthureichelberger/trailrcore/user/model"
)

type Store interface {
	CreateUser(ctx context.Context, user model.User) error
}
