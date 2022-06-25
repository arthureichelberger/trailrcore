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
)

func TestItShouldNotBeAbleToCreateUserWithoutPayload(t *testing.T) {
	cuh := handler.CreateUserHandler(service.SignupService{})
	r := gin.New()
	r.POST("/", cuh)
	req, _ := http.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestItShouldNotBeAbleToCreateUserWithWrongPayload(t *testing.T) {
	cuh := handler.CreateUserHandler(service.SignupService{})
	r := gin.New()
	r.POST("/", cuh)
	payloads := []map[string]string{
		{"email": "", "password": "", "password_confirmation": ""},
		{"email": "a", "password": "b", "password_confirmation": "c"},
	}

	for i := range payloads {
		jsonPayload, err := json.Marshal(payloads[i])
		assert.NoError(t, err)

		req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonPayload))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}

func TestItShouldNotBeAbleToCreateAUserIfAnInternalServerErrorOccurs(t *testing.T) {
	fs := store.NewFakeStore()
	fs.CreateUserHandler = func(ctx context.Context, user model.User) error {
		return fmt.Errorf("error")
	}
	sis := service.NewSignupService(fs)
	cuh := handler.CreateUserHandler(sis)
	r := gin.New()
	r.POST("/", cuh)
	jsonPayload, err := json.Marshal(map[string]string{"email": "a", "password": "b", "password_confirmation": "b"})
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonPayload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestItShouldBeAbleToCreateAUser(t *testing.T) {
	fs := store.NewFakeStore()
	sis := service.NewSignupService(fs)
	cuh := handler.CreateUserHandler(sis)
	r := gin.New()
	r.POST("/", cuh)
	jsonPayload, err := json.Marshal(map[string]string{"email": "a", "password": "b", "password_confirmation": "b"})
	assert.NoError(t, err)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonPayload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
