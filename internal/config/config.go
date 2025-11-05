package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	PostgresDSN   string
	RedisAddr     string
}

func Load() Config {
	_ = godotenv.Load("../../.env")
	fmt.Println(os.Getenv("SERVER_ADDR"))
	fmt.Println(os.Getenv("POSTGRES_DSN"))
	fmt.Println(os.Getenv("REDIS_ADDR"))
	return Config{
		os.Getenv("SERVER_ADDR"),
		os.Getenv("POSTGRES_DSN"),
		os.Getenv("REDIS_ADDR"),
	}
}
