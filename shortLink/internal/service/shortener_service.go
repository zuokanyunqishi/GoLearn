package service

import (
	"log/slog"
	"shortLink/internal/store"
)

type ShortenerService struct{}

func NewShortenerService(store store.Store, log *slog.Logger) *ShortenerService {
	return &ShortenerService{}
}
