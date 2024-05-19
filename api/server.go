package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {

	server := &Server{}
	server.setupRouter()
	return server, nil
}
func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST("/create-url", s.CreateShortUrl)
	router.GET("/:short-url", s.RetriveUrl)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s.router = router

}
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
