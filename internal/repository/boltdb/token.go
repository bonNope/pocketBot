package boltdb

import (
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

type TokenRepository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Save(chatID []byte, token []byte, bucket []byte) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		return b.Put(chatID, token)
	})
}

func (r *TokenRepository) Get(chatID []byte, bucket []byte) (string, error) {
	var token string

	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		token = string(b.Get(chatID))
		return nil
	})
	if err != nil {
		return "", err
	}

	if token == "" {
		return "", errors.New("token not found")
	}

	return token, nil
}
