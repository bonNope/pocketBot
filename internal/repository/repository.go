package repository

import (
	"github.com/boltdb/bolt"
	"github.com/bonNope/pocketBot/internal/repository/boltdb"
)

type Token interface {
	Save(chatID []byte, token []byte, bucket []byte) error
	Get(chatID []byte, bucket []byte) (string, error)
}

type Repository struct {
	Token
}

func NewRepository(db *bolt.DB) *Repository {
	return &Repository{Token: boltdb.NewTokenRepository(db)}
}
