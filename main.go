package main

import (
	"github.com/MsN-12/url_shortener/handler"
	"github.com/MsN-12/url_shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/create-url", handler.CreateShortUrl)
	r.GET("/:short-url", handler.HandleShortUrlRedirect)
	store.InitializeStore()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
