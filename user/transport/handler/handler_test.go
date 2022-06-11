//go:build integration
// +build integration

package handler_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)
	exit := m.Run()
	gin.SetMode(gin.DebugMode)

	os.Exit(exit)
}
