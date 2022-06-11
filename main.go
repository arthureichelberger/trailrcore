package main

import (
	"context"
	"fmt"

	"github.com/arthureichelberger/trailrcore/pkg/env"
	"github.com/arthureichelberger/trailrcore/pkg/pgsql"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("TrailrCore")

	db := pgsql.Connect(
		ctx,
		env.Get("TRAILRCORE_DATABASE_HOST", "localhost"),
		env.Get("TRAILRCORE_DATABASE_PORT", "5432"),
		env.Get("TRAILRCORE_DATABASE_USERNAME", "trailrcore"),
		env.Get("TRAILRCORE_DATABASE_PASSWORD", "trailrcore"),
		env.Get("TRAILRCORE_DATABASE_DB", "trailrcore"),
	)
	defer db.Close()
}
