//go:build integration
// +build integration

package http_test

import (
	"fmt"
	stdHttp "net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/arthureichelberger/trailrcore/pkg/http"
	"github.com/arthureichelberger/trailrcore/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestItShouldNotBeAbleToGoThroughAuthMiddlewareWithoutToken(t *testing.T) {
	authMiddleware := http.AuthMiddleware()
	r := gin.New()
	r.POST("/", authMiddleware)

	req, _ := stdHttp.NewRequest("POST", "/", stdHttp.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, stdHttp.StatusUnauthorized, w.Code)
}

func TestItShouldNotBeAbleToGoThroughAuthMiddlewareWithAnExpiredToken(t *testing.T) {
	authMiddleware := http.AuthMiddleware()
	r := gin.New()
	r.POST("/", authMiddleware)

	token, err := jwt.New(jwt.CustomClaims{ExpiresAt: time.Now().Add(-time.Hour).UTC().Unix()})
	assert.NoError(t, err)
	req, _ := stdHttp.NewRequest("POST", "/", stdHttp.NoBody)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, stdHttp.StatusUnauthorized, w.Code)
}

func TestItShouldBeAbleToGoThroughAuthMiddlewareWithAValidToken(t *testing.T) {
	authMiddleware := http.AuthMiddleware()
	r := gin.New()
	r.POST("/", authMiddleware)

	token, err := jwt.New(jwt.CustomClaims{ExpiresAt: time.Now().Add(time.Hour * 2).UTC().Unix()})
	assert.NoError(t, err)
	req, _ := stdHttp.NewRequest("POST", "/", stdHttp.NoBody)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, stdHttp.StatusOK, w.Code)
}
