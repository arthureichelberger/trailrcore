//go:build unit
// +build unit

package env_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/arthureichelberger/trailrcore/pkg/env"
	"github.com/stretchr/testify/assert"
)

func TestItShouldBeAbleToGetEnvValueFromEnvKey(t *testing.T) {
	randValue := func() string {
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, 6)
		rand.Read(b)
		return fmt.Sprintf("%x", b)[:6]
	}()
	os.Setenv("key", randValue)

	value := env.Get("key", "fallback")
	assert.Equal(t, value, randValue)
}

func TestItShouldBeAbleToGetEnvFallbackIfEnvKeyIsNotDefined(t *testing.T) {
	value := env.Get("undefined", "fallback")
	assert.Equal(t, "fallback", value)
}
