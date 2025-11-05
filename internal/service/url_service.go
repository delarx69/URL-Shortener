package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
	"urlShortCut/internal/db"
)

type URLService struct {
	redis *db.RedisRepo
}

func NewURLService(r *db.RedisRepo) *URLService {
	return &URLService{redis: r}
}

func generateShortCode() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (s *URLService) ShortenURL(ctx context.Context, OriginalURL string) (string, error) {
	short := generateShortCode()

	err := s.redis.Set(ctx, short, OriginalURL, 24*time.Hour)
	fmt.Printf("errr : %v", err)
	if err != nil {
		return "", err
	}

	return short, nil
}

func (s *URLService) ResolveURL(ctx context.Context, short string) (string, error) {
	return s.redis.Get(ctx, short), nil
}
