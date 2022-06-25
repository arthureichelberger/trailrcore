//go:build integration
// +build integration

package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthureichelberger/trailrcore/user/model"
	"github.com/arthureichelberger/trailrcore/user/service"
	"github.com/arthureichelberger/trailrcore/user/store"
	"github.com/arthureichelberger/trailrcore/user/transport/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestItShouldNotBeAbleToSignInWithoutPayload(t *testing.T) {
	cuh := handler.LoginHandler(service.SigninService{})
	r := gin.New()
	r.POST("/", cuh)
	req, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestItShouldNotBeAbleToSignInWithWrongPayload(t *testing.T) {
	fs := store.NewFakeStore()

	fs.GetUserByEmailHandler = func(ctx context.Context, email string) (model.User, error) {
		return model.User{}, fmt.Errorf("no user found")
	}

	sis := service.NewSigninService(fs)
	cuh := handler.LoginHandler(sis)
	r := gin.New()
	r.POST("/", cuh)
	payloads := []map[string]string{
		{"email": "", "password": ""},
		{"email": "a", "password": ""},
		{"email": "a", "password": "b"},
	}

	for i := range payloads {
		jsonPayload, err := json.Marshal(payloads[i])
		assert.NoError(t, err)

		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonPayload))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}
}

func TestItShouldBeAbleToSignIn(t *testing.T) {
	fs := store.NewFakeStore()

	pwd, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := model.User{
		Email:    "a@gmail.com",
		Password: string(pwd),
	}

	fs.GetUserByEmailHandler = func(ctx context.Context, email string) (model.User, error) {
		return user, nil
	}

	sis := service.NewSigninService(fs)
	cuh := handler.LoginHandler(sis)
	r := gin.New()
	r.POST("/", cuh)

	jsonPayload, err := json.Marshal(map[string]string{"email": "a@gmail.com", "password": "test"})
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonPayload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.HeaderMap["Authorization"])
}
