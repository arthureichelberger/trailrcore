//go:build unit
// +build unit

package jwt_test

import (
	"testing"
	"time"

	"github.com/arthureichelberger/trailrcore/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestClaimsShouldNotBeValidIfExpiredAtIsAlreadyPast(t *testing.T) {
	claims := jwt.CustomClaims{ExpiresAt: time.Now().Add(-time.Hour).UnixMilli()}
	err := claims.Valid()
	assert.Error(t, err)
}

func TestClaimsShouldBeValidIfExpiresAtIsInTheFuture(t *testing.T) {
	claims := jwt.CustomClaims{ExpiresAt: time.Now().Add(time.Hour).UnixMilli()}
	err := claims.Valid()
	assert.NoError(t, err)
}
