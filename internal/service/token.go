package service

import (
	"github.com/bonNope/pocketBot/internal/repository"
	"strconv"
)

type TokenService struct {
	repo repository.Token
}

func NewTokenService(repo repository.Token) *TokenService {
	return &TokenService{repo: repo}
}

func (s *TokenService) Save(chatID int64, token string, bucket Bucket) error {
	return s.repo.Save(s.intToBytes(chatID), []byte(token), []byte(bucket))
}
func (s *TokenService) Get(chatID int64, bucket Bucket) (string, error) {
	return s.repo.Get(s.intToBytes(chatID), []byte(bucket))
}
func (s *TokenService) intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
