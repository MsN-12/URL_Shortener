package handler

import (
	"github.com/MsN-12/url_shortener/shortener"
	"github.com/MsN-12/url_shortener/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Todo: tests
type CreateShortUrlRequest struct {
	LongUrl string `json:"long_url"`
}

func CreateShortUrl(ctx *gin.Context) {
	var request CreateShortUrlRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for {
		shortUrl := shortener.GenerateShortURL(request.LongUrl)
		if store.CheckExistence(shortUrl) {
			continue
		} else {
			store.SaveUrl(shortUrl, request.LongUrl)
			ctx.JSON(http.StatusOK, gin.H{"short_url": shortUrl})
			break
		}
	}
}

type ShortUrlRedirectRequest struct {
	ShortUrl string `json:"short_url"`
}

func RetriveUrl(ctx *gin.Context) {
	var request ShortUrlRedirectRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	longUrl := store.RetriveUrl(request.ShortUrl)
	ctx.JSON(http.StatusOK, gin.H{"long_url": longUrl})
}
