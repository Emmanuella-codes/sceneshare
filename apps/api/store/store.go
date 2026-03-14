package store

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("not found")
var ErrExpired = errors.New("link expired")
var ErrForbidden = errors.New("forbidden")

type Store struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, databaseURL string) (*Store, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("creating pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pinging db: %w", err)
	}

	return &Store{db: pool}, nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) CreateLink(ctx context.Context, params dtos.CreateLinkParams) (*models.Link, error) {
	query := `
		INSERT INTO links (short_code, platform, content_id, timestamp_s, title, thumbnail, owner_token, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, short_code, platform, content_id, timestamp_s, title, thumbnail, owner_token, created_by, created_at, expires_at, click_count
	`

	row := s.db.QueryRow(ctx, query,
		params.ShortCode,
		params.Platform,
		params.ContentID,
		params.TimestampS,
		params.Title,
		params.Thumbnail,
		params.OwnerToken,
		params.ExpiresAt,
	)
	return scanLink(row)
}

func (s *Store) GetLinkByCode(ctx context.Context, code string) (*models.Link, error) {
	query := `
		SELECT id, short_code, platform, content_id, timestamp_s, title, thumbnail, owner_token, created_by, created_at, expires_at, click_count
		FROM links
		WHERE short_code = $1
	`
	row := s.db.QueryRow(ctx, query, code)
	link, err := scanLink(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("getting link: %w", err)
	}

	if link.ExpiresAt != nil && time.Now().After(*link.ExpiresAt) {
		return nil, ErrExpired
	}
	return link, nil
}

func (s *Store) DeleteLink(ctx context.Context, code, token string) error {
	result, err := s.db.Exec(ctx, `DELETE FROM links WHERE short_code = $1 AND owner_token = $2`, code, token)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		var exists bool
		if err := s.db.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM links WHERE short_code = $1)`, code).Scan(&exists); err != nil {
			return err
		}
		if !exists {
			return ErrNotFound
		}
		return ErrForbidden
	}
	return nil
}

func (s *Store) IncrementClickCount(ctx context.Context, event models.ClickEvent) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`UPDATE links SET click_count = click_count + 1 WHERE id = $1`,
		event.LinkID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		`INSERT INTO click_events (link_id, user_agent, referrer) VALUES ($1, $2, $3)`,
		event.LinkID, event.UserAgent, event.Referrer)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func scanLink(row pgx.Row) (*models.Link, error) {
	var link models.Link
	err := row.Scan(
		&link.ID,
		&link.ShortCode,
		&link.Platform,
		&link.ContentID,
		&link.TimestampS,
		&link.Title,
		&link.Thumbnail,
		&link.OwnerToken,
		&link.CreatedBy,
		&link.CreatedAt,
		&link.ExpiresAt,
		&link.ClickCount,
	)
	if err != nil {
		return nil, fmt.Errorf("scanning link: %w", err)
	}
	return &link, nil
}
