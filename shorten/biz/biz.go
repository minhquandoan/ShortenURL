package biz

import (
	"context"
	"fmt"

	"github.com/quandoan/shorten_url/modules/shorten/rand"
)

type ShortenStore interface {
	Create(ctx context.Context, url string, code string) error
}

type Hashing interface {
	Hash(string) string
}

type shortenBiz struct {
	store ShortenStore
	hash Hashing
}

func NewShortenBiz(store ShortenStore, hash Hashing) *shortenBiz {
	return &shortenBiz{
		store: store,
		hash: hash,
	}
}

func(biz *shortenBiz) CreateVirtualLink(ctx context.Context, url string) (string, error) {
	salt := rand.GenSalt(10)
	code := biz.hash.Hash(salt)
	err := biz.store.Create(ctx, url, code)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("localhost:8080/%s", code), err
}
