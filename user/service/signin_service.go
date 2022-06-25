package service

import (
	"context"
	"fmt"
	"time"

	"github.com/arthureichelberger/trailrcore/pkg/jwt"
	"github.com/arthureichelberger/trailrcore/user/store"
	"golang.org/x/crypto/bcrypt"
)

type SigninService struct {
	userStore store.Store
}

func NewSigninService(userStore store.Store) SigninService {
	return SigninService{
		userStore: userStore,
	}
}

func (sis SigninService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := sis.userStore.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("no user found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("wrong password")
	}

	expire := time.Now().Add(time.Hour * 2)
	token, err := jwt.New(jwt.CustomClaims{
		ExpiresAt: expire.UnixMilli(),
		Claims: map[string]interface{}{
			"id":    user.ID,
			"email": email,
		},
	})

	return token, err
}
