package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/store"
)

func TestDeleteLink(t *testing.T) {
	tests := []struct {
		name        string
		code        string
		token       string
		store       *mockStore
		expectedErr error
	}{
		{
			name:  "success",
			code:  "abc1234",
			token: "owner-token",
			store: &mockStore{
				deleteLink: func(_ context.Context, code, token string) error {
					if code != "abc1234" {
						t.Fatalf("expected code %q, got %q", "abc1234", code)
					}
					if token != "owner-token" {
						t.Fatalf("expected token %q, got %q", "owner-token", token)
					}
					return nil
				},
			},
		},
		{
			name:  "not found",
			code:  "missing1",
			token: "owner-token",
			store: &mockStore{
				deleteLink: func(context.Context, string, string) error {
					return store.ErrNotFound
				},
			},
			expectedErr: ErrNotFound,
		},
		{
			name:  "forbidden",
			code:  "abc1234",
			token: "wrong-token",
			store: &mockStore{
				deleteLink: func(context.Context, string, string) error {
					return store.ErrForbidden
				},
			},
			expectedErr: ErrForbidden,
		},
		{
			name:  "internal error",
			code:  "broken12",
			token: "owner-token",
			store: &mockStore{
				deleteLink: func(context.Context, string, string) error {
					return errors.New("boom")
				},
			},
			expectedErr: errors.New("boom"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewLinkService(tc.store, "http://localhost:3001")
			err := svc.DeleteLink(context.Background(), tc.code, tc.token)

			if tc.expectedErr == nil {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}

			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if err.Error() != tc.expectedErr.Error() {
				t.Fatalf("expected error %q, got %q", tc.expectedErr.Error(), err.Error())
			}
		})
	}
}
