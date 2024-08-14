package handlers

import (
	"github.com/gin-contrib/cors"
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
	router.Use(cors.Default())

	router.POST("/api/create-url", s.CreateShortUrl)
	router.GET("/:short-url", s.RetrieveUrl)

	s.router = router

}
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
