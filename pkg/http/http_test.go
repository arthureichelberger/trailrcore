//go:build integration
// +build integration

package http_test

import (
	"context"
	"fmt"
	stdHttp "net/http"
	"sync"
	"testing"
	"time"

	"github.com/arthureichelberger/trailrcore/pkg/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestItShouldBeAbleToServeHTTPServer(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := stdHttp.HandlerFunc(func(w stdHttp.ResponseWriter, r *stdHttp.Request) {
		w.WriteHeader(stdHttp.StatusOK)
	})

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		assert.NoError(t, http.Serve(ctx, "127.0.0.1:0", handler))
	}()

	cancel()
	wg.Wait()
}

func TestHTTPRecovery(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(http.Recovery())
	g.NoRoute(func(ctx *gin.Context) {
		panic("test")
	})

	wg := new(sync.WaitGroup)
	wg.Add(1)
	addr := "127.0.0.1:8001"
	go func() {
		defer wg.Done()
		err := http.Serve(ctx, addr, g)
		assert.NoError(t, err)
	}()

	time.Sleep(time.Millisecond * 100)
	res, err := stdHttp.DefaultClient.Get(fmt.Sprintf("http://%s%s", addr, "/"))
	if err != nil {
		t.FailNow()
	}

	defer res.Body.Close()
	assert.NotNil(t, res)
	assert.Equal(t, stdHttp.StatusInternalServerError, res.StatusCode)

	cancel()
	wg.Wait()
}
