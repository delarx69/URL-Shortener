package main

import (
	"fmt"
	"urlShortCut/internal/config"
	"urlShortCut/internal/db"
	"urlShortCut/internal/handler"
	"urlShortCut/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	redisClient := db.InitRedis(cfg)
	urlSvc := service.NewURLService(redisClient)
	fmt.Printf("redisClinet: %T", redisClient)

	r := gin.Default()
	r.POST("/reg", handler.AddShortedURL(urlSvc))

	r.GET("/resolve/:short", handler.GetOriginalURL(urlSvc))

	if err := r.Run(cfg.ServerAddress); err != nil {
		fmt.Printf("Не удалось запустить сервер: %v\n", err)
	}
}
