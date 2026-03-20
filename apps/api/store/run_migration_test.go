package store

import (
	"context"
	"testing"
)

func TestRunMigrations_Idempotent(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()

	if err := s.RunMigrations(ctx); err != nil {
		t.Fatalf("rerun migrations: %v", err)
	}

	var count int
	if err := s.db.QueryRow(ctx, `SELECT COUNT(*) FROM schema_migrations`).Scan(&count); err != nil {
		t.Fatalf("count schema migrations: %v", err)
	}
	if count < 2 {
		t.Fatalf("expected at least %d applied migrations, got %d", 2, count)
	}
}
