package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/go-chi/chi/v5"
)

type mockLinkService struct {
	createLink         func(context.Context, *dtos.CreateLinkInput) (*dtos.LinkResponse, error)
	getLink            func(context.Context, string) (*dtos.LinkResponse, error)
	getLinkForRedirect func(context.Context, string) (*models.Link, error)
	recordClick        func(string, string, string)
	getStats           func(context.Context, string) (*dtos.StatsResponse, error)
	deleteLink         func(context.Context, string, string) error
}

func (m *mockLinkService) CreateLink(ctx context.Context, input *dtos.CreateLinkInput) (*dtos.LinkResponse, error) {
	if m.createLink == nil {
		panic("mockLinkService.createLink is not set")
	}
	return m.createLink(ctx, input)
}

func (m *mockLinkService) GetLink(ctx context.Context, code string) (*dtos.LinkResponse, error) {
	if m.getLink == nil {
		panic("mockLinkService.getLink is not set")
	}
	return m.getLink(ctx, code)
}

func (m *mockLinkService) GetLinkForRedirect(ctx context.Context, code string) (*models.Link, error) {
	if m.getLinkForRedirect == nil {
		panic("mockLinkService.getLinkForRedirect is not set")
	}
	return m.getLinkForRedirect(ctx, code)
}

func (m *mockLinkService) RecordClick(id string, userAgent, referer string) {
	if m.recordClick != nil {
		m.recordClick(id, userAgent, referer)
	}
}

func (m *mockLinkService) GetStats(ctx context.Context, code string) (*dtos.StatsResponse, error) {
	if m.getStats == nil {
		panic("mockLinkService.getStats is not set")
	}
	return m.getStats(ctx, code)
}

func (m *mockLinkService) DeleteLink(ctx context.Context, code, token string) error {
	if m.deleteLink == nil {
		panic("mockLinkService.deleteLink is not set")
	}
	return m.deleteLink(ctx, code, token)
}

func withCode(r *http.Request, code string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("code", code)
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func decodeJSON[T any](t *testing.T, resp *http.Response) T {
	t.Helper()

	var v T
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	return v
}
