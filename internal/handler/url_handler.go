package handler

import (
	"encoding/base64"
	"net/http"
	"urlShortCut/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

type URLRequest struct {
	OriginalURL string `json:"OriginalURL"`
}

type URLResponse struct {
	Short    string `json:"short"`
	Original string `json:"original"`
}

func AddShortedURL(svc *service.URLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req URLRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx := c.Request.Context()
		short, err := svc.ShortenURL(ctx, req.OriginalURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		shortURL := "http://localhost:8080/resolve/" + short

		qr, err := qrcode.Encode(shortURL, qrcode.Medium, 256)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "QR generation failed"})
			return
		}

		qrBase64 := base64.StdEncoding.EncodeToString(qr)

		c.JSON(http.StatusOK, gin.H{
			"short_url":    shortURL,
			"original_url": req.OriginalURL,
			"qr_code":      "data:image/png;base64," + qrBase64,
		})
	}
}

func GetOriginalURL(svc *service.URLService) gin.HandlerFunc {
	return func(c *gin.Context) {
		short := c.Param("short")
		ctx := c.Request.Context()
		original, err := svc.ResolveURL(ctx, short)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.Redirect(http.StatusFound, original)

	}
}
