package service

import (
	"context"
	"fmt"
	"time"

	"github.com/arthureichelberger/trailrcore/user/exception"
	"github.com/arthureichelberger/trailrcore/user/model"
	"github.com/arthureichelberger/trailrcore/user/store"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type SignInService struct {
	userStore store.Store
}

func NewSignInService(userStore store.Store) SignInService {
	return SignInService{
		userStore: userStore,
	}
}

func (sis SignInService) CreateUser(ctx context.Context, email, pwd, pwdConfirm string) (model.User, error) {
	switch {
	case email == "":
		return model.User{}, exception.CouldNotValidateEmailError{}
	case pwd != pwdConfirm, pwd == "":
		return model.User{}, exception.CouldNotValidatePasswordError{}
	}

	bcryptPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("could not generate bcrypt password")
		return model.User{}, fmt.Errorf("could not generate password")
	}

	user := model.User{
		ID:        uuid.New(),
		Email:     email,
		Password:  string(bcryptPwd),
		CreatedAt: time.Now(),
	}

	if err := sis.userStore.CreateUser(ctx, user); err != nil {
		log.Error().Err(err).Str("email", user.Email).Msg("could not create user")
		return model.User{}, fmt.Errorf("could not create user")
	}

	return user, nil
}
