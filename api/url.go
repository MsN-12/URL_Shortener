package api

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

func (s *Server) CreateShortUrl(ctx *gin.Context) {
	var request CreateShortUrlRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	existingShortUrl, err := store.RetrieveUrl(request.LongUrl)
	if err != nil {
		if err.Error() == "redis: nil" {
			for {
				shortUrl := shortener.GenerateShortURL(request.LongUrl)
				exists, err := store.CheckExistence(shortUrl)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, errorResponse(err))
					return
				}
				if exists {
					continue
				} else {
					err := store.SaveUrl(shortUrl, request.LongUrl)
					if err != nil {
						ctx.JSON(http.StatusBadRequest, errorResponse(err))
						return
					}
					ctx.JSON(http.StatusOK, gin.H{"short_url": shortUrl})
					return
				}
			}
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"short_url": existingShortUrl})
	return
}

func (s *Server) RetriveUrl(ctx *gin.Context) {
	shortUrl, ok := ctx.Params.Get("short-url")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": "invalid url"})
		return
	}
	longUrl, err := store.RetrieveUrl(shortUrl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"long_url": longUrl})
}
