package handlers

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"net/http"
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

	router.StaticFS("/public/", http.Dir("public"))
	router.POST("/api/create-url", s.CreateShortUrl)
	router.GET("/api/:short-url", s.RetrieveUrl)

	router.GET("/", s.Home)
	router.GET("/auth/singup", s.Auth)

	s.router = router

}
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func Render(ctx *gin.Context, c templ.Component) {
	err := c.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, errorResponse(err))
		return
	}
}
