package main

import (
	"context"
	"fmt"
	stdHttp "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/arthureichelberger/trailrcore/pkg/env"
	"github.com/arthureichelberger/trailrcore/pkg/http"
	"github.com/arthureichelberger/trailrcore/pkg/pgsql"
	"github.com/arthureichelberger/trailrcore/user/service"
	"github.com/arthureichelberger/trailrcore/user/store"
	"github.com/arthureichelberger/trailrcore/user/transport"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := pgsql.Connect(
		ctx,
		env.Get("TRAILRCORE_DATABASE_HOST", "localhost"),
		env.Get("TRAILRCORE_DATABASE_PORT", "5432"),
		env.Get("TRAILRCORE_DATABASE_USERNAME", "trailrcore"),
		env.Get("TRAILRCORE_DATABASE_PASSWORD", "trailrcore"),
		env.Get("TRAILRCORE_DATABASE_DB", "trailrcore"),
	)
	defer db.Close()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		defer cancel()
		<-sigc
	}()

	errgrp, ctx := errgroup.WithContext(ctx)

	errgrp.Go(func() error {
		g := gin.New()
		g.GET(env.Get("TRAILRCORE_HTTP_HEARTBEAT", "/trailrcore/__internal__/heartbeat"), func(ctx *gin.Context) {
			ctx.JSON(stdHttp.StatusOK, gin.H{"ping": "pong"})
		})

		mainGrp := g.Group("/trailrcore")
		initRouters(ctx, mainGrp, db)

		return http.Serve(
			ctx,
			fmt.Sprintf("%s:%s", env.Get("TRAILRCORE_HTTP_HOST", "0.0.0.0"), env.Get("TRAILRCORE_HTTP_PORT", "8080")),
			g,
		)
	})

	if err := errgrp.Wait(); err != nil {
		log.Panic().Err(err).Msg("service shutdown")
	}
}

func initRouters(_ context.Context, engine *gin.RouterGroup, db *sqlx.DB) {
	userStore := store.NewPgStore(db)
	signInService := service.NewSignInService(userStore)

	userRouter := transport.NewUserRouter(signInService)
	userRouter.InitRouter(engine)
}
