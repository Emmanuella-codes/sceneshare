package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Emmanuella-codes/sceneshare/api/config"
	"github.com/Emmanuella-codes/sceneshare/api/store"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()

	db, err := store.New(ctx, cfg.DatabaseURL)
	if err != nil {
		slog.Error("connecting to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("connected to database")

	// 
}
