package handler

import (
	"log/slog"
	"net/http"
	"shortLink/internal/service"
)

type ShortenerHandler struct {
}

func NewLinkHandler(svc *service.ShortenerService, with *slog.Logger) *ShortenerHandler {

	return &ShortenerHandler{}
}

func (h *ShortenerHandler) CreateShortLink(w http.ResponseWriter, r *http.Request) {

}

func (h *ShortenerHandler) RedirectShortLink(w http.ResponseWriter, r *http.Request, code string) {

}
