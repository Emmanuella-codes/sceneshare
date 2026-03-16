package store

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/Emmanuella-codes/sceneshare/api/migrations"
)

// applies each migration file once and records the applied versions.
func (s *Store) RunMigrations(ctx context.Context) error {
	if _, err := s.db.Exec(ctx, `CREATE EXTENSION IF NOT EXISTS pgcrypto`); err != nil {
		return fmt.Errorf("creating pgcrypto extension: %w", err)
	}

	if _, err := s.db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`); err != nil {
		return fmt.Errorf("creating schema_migrations table: %w", err)
	}

	files, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return fmt.Errorf("loading migration files: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		applied, err := s.isMigrationApplied(ctx, file.Name())
		if err != nil {
			return err
		}
		if applied {
			continue
		}

		sqlBytes, err := migrations.Files.ReadFile(file.Name())
		if err != nil {
			return fmt.Errorf("reading migration %s: %w", file.Name(), err)
		}

		if err := s.applyMigration(ctx, file.Name(), string(sqlBytes)); err != nil {
			return err
		}
	}

	return nil
}

func (s *Store) isMigrationApplied(ctx context.Context, version string) (bool, error) {
	var applied bool
	if err := s.db.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)`, version).Scan(&applied); err != nil {
		return false, fmt.Errorf("checking migration %s: %w", version, err)
	}
	return applied, nil
}

// wraps each migration in a transaction so version tracking stays consistent.
func (s *Store) applyMigration(ctx context.Context, version, sql string) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("starting migration %s: %w", version, err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, sql); err != nil {
		return fmt.Errorf("applying migration %s: %w", version, err)
	}

	if _, err := tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, version); err != nil {
		return fmt.Errorf("recording migration %s: %w", version, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("committing migration %s: %w", version, err)
	}

	return nil
}
