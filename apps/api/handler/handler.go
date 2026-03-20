package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/Emmanuella-codes/sceneshare/api/service"
	"github.com/Emmanuella-codes/sceneshare/api/utils"
	"github.com/go-chi/chi/v5"
)

type linkService interface {
	CreateLink(ctx context.Context, input *dtos.CreateLinkInput) (*dtos.LinkResponse, error)
	GetLink(ctx context.Context, code string) (*dtos.LinkResponse, error)
	DeleteLink(ctx context.Context, code, token string) error
	GetStats(ctx context.Context, code string) (*dtos.StatsResponse, error)
	GetLinkForRedirect(ctx context.Context, code string) (*models.Link, error)
	RecordClick(id string, userAgent, referer string)
}

type Handler struct {
	links linkService
}

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func New(links linkService) *Handler {
	return &Handler{links: links}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var input dtos.CreateLinkInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "request body is not valid JSON")
		return
	}

	resp, err := h.links.CreateLink(r.Context(), &input)
	if err != nil {
		var ve *utils.ValidationError
		if errors.As(err, &ve) {
			writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", ve.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "failed to create link")
		return
	}
	writeJSON(w, http.StatusCreated, resp)
}

func (h *Handler) GetLink(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	resp, err := h.links.GetLink(r.Context(), code)
	if err != nil {
		writeLinkError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	token := r.Header.Get("X-Owner-Token")
	if token == "" {
		writeError(w, http.StatusUnauthorized, "MISSING_TOKEN", "X-Owner-Token header is required")
		return
	}
	if err := h.links.DeleteLink(r.Context(), code, token); err != nil {
		writeLinkError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	resp, err := h.links.GetStats(r.Context(), code)
	if err != nil {
		writeLinkError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	link, err := h.links.GetLinkForRedirect(r.Context(), code)
	if err != nil {
		writeLinkError(w, err)
		return
	}

	go h.links.RecordClick(link.ID, r.UserAgent(), r.Referer())

	target, err := service.BuildDeepLink(link)
	if err != nil {
		writeLinkError(w, err)
		return
	}
	http.Redirect(w, r, target, http.StatusFound)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, errorResponse{Code: code, Message: message})
}

func writeLinkError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrNotFound):
		writeError(w, http.StatusNotFound, "LINK_NOT_FOUND", "link does not exist")
	case errors.Is(err, service.ErrExpired):
		writeError(w, http.StatusGone, "LINK_EXPIRED", "link has expired")
	case errors.Is(err, service.ErrForbidden):
		writeError(w, http.StatusForbidden, "FORBIDDEN", "invalid owner token")
	case errors.Is(err, service.ErrUnsupportedPlatform):
		writeError(w, http.StatusInternalServerError, "UNSUPPORTED_PLATFORM", "link platform is not configured for redirects")
	default:
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "something went wrong")
	}
}
