package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cleidison-barradas/shortr.api/internal/config"
	"github.com/cleidison-barradas/shortr.api/internal/infra/postgres"
	"github.com/cleidison-barradas/shortr.api/internal/infra/server"
	"github.com/cleidison-barradas/shortr.api/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := run(); err != nil {
		slog.Error("application error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	dbCtx, cancelDB := context.WithTimeout(ctx, 5*time.Second)
	pool, err := postgres.NewPool(dbCtx, cfg)
	cancelDB()
	if err != nil {
		return fmt.Errorf("create pool: %w", err)
	}
	defer pool.Close()

	srv := server.New(cfg, router.New(pool, cfg))

	serverErr := make(chan error, 1)
	go func() {
		serverErr <- srv.Listen()
	}()

	select {
	case err := <-serverErr:
		return fmt.Errorf("http server: %w", err)
	case <-ctx.Done():
		slog.Info("shutdown signal received")
	}

	stop()

	shutDownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutDownCtx); err != nil {
		return fmt.Errorf("shutdown http server: %w", err)
	}

	slog.Info("server shutdown gracefully")
	return nil
}
