package store

import (
	"log/slog"
)

type Store interface {
	get(key string) string
}

// type Store struct{}
type Storer struct {
}

func (s Storer) get(key string) string {
	return ""

}
func NewStore(storeType string, log *slog.Logger) Store {
	return Storer{}
}
