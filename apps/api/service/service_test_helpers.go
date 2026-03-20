package service

import (
	"context"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
)

type mockStore struct {
	createLink          func(context.Context, dtos.CreateLinkParams) (*models.Link, error)
	getLinkByCode       func(context.Context, string) (*models.Link, error)
	deleteLink          func(context.Context, string, string) error
	incrementClickCount func(context.Context, models.ClickEvent) error
}

func (m *mockStore) CreateLink(ctx context.Context, params dtos.CreateLinkParams) (*models.Link, error) {
	if m.createLink == nil {
		panic("mockStore.CreateLink not configured")
	}
	return m.createLink(ctx, params)
}

func (m *mockStore) GetLinkByCode(ctx context.Context, code string) (*models.Link, error) {
	if m.getLinkByCode == nil {
		panic("mockStore.GetLinkByCode not configured")
	}
	return m.getLinkByCode(ctx, code)
}

func (m *mockStore) DeleteLink(ctx context.Context, code, token string) error {
	if m.deleteLink == nil {
		panic("mockStore.DeleteLink not configured")
	}
	return m.deleteLink(ctx, code, token)
}

func (m *mockStore) IncrementClickCount(ctx context.Context, event models.ClickEvent) error {
	if m.incrementClickCount == nil {
		panic("mockStore.IncrementClickCount not configured")
	}
	return m.incrementClickCount(ctx, event)
}
