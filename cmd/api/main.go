package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Je33/packager/internal/config"
	"github.com/Je33/packager/internal/repository/mem"
	"github.com/Je33/packager/internal/service/packer"
	"github.com/Je33/packager/internal/transport/graphql"
	"github.com/Je33/packager/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	log := logger.New(logger.Config{
		Level: cfg.Log.Level,
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	repo := mem.New(log)

	pack := packer.New(repo, log)

	server := graphql.New(cfg.GraphQL, pack, log)

	log.Info("Starting server on :" + cfg.GraphQL.Port)
	err = server.Start(ctx)
	if err != nil {
		log.Error("Server stopped", "error", err)
		return
	}

	// TODO: add graceful shutdown

	log.Info("Server stopped")
}
