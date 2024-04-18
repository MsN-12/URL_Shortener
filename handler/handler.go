package handler

import (
	"github.com/MsN-12/url_shortener/shortener"
	"github.com/MsN-12/url_shortener/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
func HandleShortUrlRedirect(ctx *gin.Context) {
	shortUrl := ctx.Param("short-url")
	originalUrl := store.RetriveUrl(shortUrl)
	ctx.Redirect(http.StatusFound, originalUrl)
}
