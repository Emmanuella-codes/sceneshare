package store

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func testDatabaseURL(t *testing.T) string {
	t.Helper()

	if v := os.Getenv("TEST_DATABASE_URL"); v != "" {
		return v
	}

	envPath := filepath.Join("..", ".env")
	values, err := godotenv.Read(envPath)
	if err != nil {
		t.Fatalf("read %s: %v", envPath, err)
	}

	v := values["TEST_DATABASE_URL"]
	if v == "" {
		t.Fatal("TEST_DATABASE_URL is required")
	}

	return v
}

func newTestStore(t *testing.T) *Store {
	t.Helper()

	ctx := context.Background()
	s, err := New(ctx, testDatabaseURL(t))
	if err != nil {
		t.Fatalf("new test store: %v", err)
	}

	if err := s.RunMigrations(ctx); err != nil {
		s.Close()
		t.Fatalf("run migrations: %v", err)
	}

	clearTestData(t, s)

	t.Cleanup(func() {
		clearTestData(t, s)
		s.Close()
	})

	return s
}

func clearTestData(t *testing.T, s *Store) {
	t.Helper()

	if _, err := s.db.Exec(context.Background(), `TRUNCATE TABLE click_events, links RESTART IDENTITY CASCADE`); err != nil {
		t.Fatalf("clear test data: %v", err)
	}
}
