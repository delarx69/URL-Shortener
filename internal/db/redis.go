package db

import (
	"context"
	"log"
	"time"
	"urlShortCut/internal/config"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	client *redis.Client
}

func NewRedisRepo(c *redis.Client) *RedisRepo {
	return &RedisRepo{client: c}
}

func (r *RedisRepo) Get(ctx context.Context, key string) string {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (r *RedisRepo) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

func InitRedis(cfg config.Config) *RedisRepo {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("Ошибка подключения к Redis: %v", err, "\n")
	}

	log.Println("Подключение к Redis успешно")
	return NewRedisRepo(rdb)
}
