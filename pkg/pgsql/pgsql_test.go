//go:build integration
// +build integration

package pgsql_test

import (
	"context"
	"testing"

	"github.com/arthureichelberger/trailrcore/pkg/pgsql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestItShouldNotBeAbleToConnectToANonExistingHost(t *testing.T) {
	assert.Panics(t, func() {
		pgsql.Connect(context.Background(), "test", "5432", "", "", "")
	})
}

func TestItShouldNotBeAbleToConnectWithIncorrectCredentials(t *testing.T) {
	assert.Panics(t, func() {
		pgsql.Connect(context.Background(), "127.0.0.1", "5432", "c", "c", "c")
	})
}

func TestItShouldBeAbleToConnectWithCorrectCredentials(t *testing.T) {
	assert.NotPanics(t, func() {
		db := pgsql.Connect(context.Background(), "127.0.0.1", "5432", "trailrcore", "trailrcore", "trailrcore")
		assert.IsType(t, new(sqlx.DB), db)
	})
}
