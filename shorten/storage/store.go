package storage

import (
	"context"
	"errors"

	"github.com/quandoan/shorten_url/db"
)

type store struct {
	// collection *mongo.Collection
	db db.MemDb
}

func NewShortenStore(db db.MemDb) *store {
	return &store{db: db}
}

func(s *store) Create(ctx context.Context, url string, code string) error {
	if _, ok := s.db.Store[code]; ok {
		return errors.New("Already existing")
	}
	
	s.db.Store[code] = url
	return nil
}
