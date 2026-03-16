package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/config"
	"github.com/Emmanuella-codes/sceneshare/api/handler"
	"github.com/Emmanuella-codes/sceneshare/api/pkg"
	"github.com/Emmanuella-codes/sceneshare/api/service"
	"github.com/Emmanuella-codes/sceneshare/api/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
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

	// Apply the current schema on startup so local development does not depend on manual setup.
	if err := db.RunMigrations(ctx); err != nil {
		slog.Error("running migrations", "error", err)
		os.Exit(1)
	}
	slog.Info("database schema ready")

	linkService := service.NewLinkService(db, cfg.BaseURL)
	h := handler.New(linkService)

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(pkg.Logger)
	r.Use(httprate.LimitByIP(100, time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: cfg.AllowedOrigins,
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "X-Owner-Token"},
	}))

	// routes
	r.Get("/health", h.Health)
	r.Get("/r/{code}", h.Redirect)

	r.Route("/api/v1", func(r chi.Router) {
		r.With(httprate.LimitByIP(10, time.Minute)).Post("/links", h.CreateLink)
		r.Get("/links/{code}", h.GetLink)
		r.Delete("/links/{code}", h.DeleteLink)
		r.Get("/links/{code}/stats", h.GetStats)
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		slog.Info("server starting", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	// block until we receive SIGINT or SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "error", err)
	}
	slog.Info("stopped")
}
