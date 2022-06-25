//go:build unit
// +build unit

package jwt_test

import (
	"testing"

	"github.com/arthureichelberger/trailrcore/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestItShouldBeAbleToCreateJWT(t *testing.T) {
	token, err := jwt.New(jwt.CustomClaims{})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
