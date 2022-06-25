//go:build unit
// +build unit

package jwt_test

import (
	"testing"
	"time"

	"github.com/arthureichelberger/trailrcore/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestItShouldBeAbleToCreateJWT(t *testing.T) {
	token, err := jwt.New(jwt.CustomClaims{})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestItShouldNotBeAbleToDecodeAnExpiredJWT(t *testing.T) {
	token, err := jwt.New(jwt.CustomClaims{
		ExpiresAt: time.Now().Add(-time.Hour).UTC().Unix(),
	})
	assert.NoError(t, err)

	claims, err := jwt.Decode(token)
	assert.Error(t, err)
	assert.Empty(t, claims)
}

func TestItShouldBeAbleToDecodeJWT(t *testing.T) {
	token, err := jwt.New(jwt.CustomClaims{
		ExpiresAt: time.Now().Add(time.Hour * 2).UTC().Unix(),
	})
	assert.NoError(t, err)

	claims, err := jwt.Decode(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, claims)
}
