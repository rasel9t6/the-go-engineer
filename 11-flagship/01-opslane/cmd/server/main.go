// Copyright (c) 2026 Rasel Hossen
// See LICENSE for usage terms.

package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/rasel9t6/the-go-engineer/11-flagship/01-opslane/internal/config"
	"github.com/rasel9t6/the-go-engineer/11-flagship/01-opslane/internal/db"
	"github.com/rasel9t6/the-go-engineer/11-flagship/01-opslane/internal/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.New(slog.NewTextHandler(os.Stderr, nil)).Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.App.LogLevel,
	}))

	ctx := context.Background()
	database, err := db.Open(ctx, cfg.Database)
	if err != nil {
		logger.Error("failed to open database", slog.Any("error", err))
		os.Exit(1)
	}
	defer database.Close()

	if err := db.Migrate(ctx, database); err != nil {
		logger.Error("failed to apply database migrations", slog.Any("error", err))
		os.Exit(1)
	}

	store := db.NewStore(database)
	app := &handlers.Application{
		Logger:      logger,
		Store:       store,
		ServiceName: cfg.App.Name,
		Environment: cfg.App.Env,
	}

	server := &http.Server{
		Addr:              cfg.HTTP.Address,
		Handler:           app.Routes(),
		ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
		ReadTimeout:       cfg.HTTP.ReadTimeout,
		WriteTimeout:      cfg.HTTP.WriteTimeout,
		IdleTimeout:       cfg.HTTP.IdleTimeout,
	}

	logger.Info("starting opslane server",
		slog.String("env", cfg.App.Env),
		slog.String("addr", cfg.HTTP.Address),
		slog.String("database", "postgresql"),
	)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server stopped unexpectedly", slog.Any("error", err))
		os.Exit(1)
	}
}
