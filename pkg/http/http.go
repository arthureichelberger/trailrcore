// Package http contains generic primitives to create an HTTP server.
package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		log.Ctx(c.Request.Context()).Error().Interface("error", err).Msg("panic recovery")
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}

func Serve(ctx context.Context, addr string, handler http.Handler) error {
	srv := new(http.Server)
	srv.Addr = addr
	srv.Handler = handler

	sink := make(chan error, 1)

	go func() {
		defer close(sink)
		sink <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		return shutdown(srv)
	case err := <-sink:
		return err
	}
}

const shutdownTimeout = 10 * time.Second

func shutdown(srv *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctx)
}
