package service

import (
	"github.com/bonNope/pocketBot/internal/repository"
	"github.com/bonNope/pocketBot/pkg/pocket"
)

type Bucket string

const (
	AccessTokens  Bucket = "access_tokens"
	RequestTokens Bucket = "request_tokens"
)

type Token interface {
	Save(chatID int64, token string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
	intToBytes(v int64) []byte
}

type Service struct {
	Token
	PocketClient *pocket.Client
}

func NewService(repo *repository.Repository, client *pocket.Client) *Service {
	return &Service{
		Token:        NewTokenService(repo.Token),
		PocketClient: client,
	}
}
