package handlers

import (
	"github.com/MsN-12/url_shortener/view/auth"
	"github.com/MsN-12/url_shortener/view/home"
	"github.com/gin-gonic/gin"
)

func (s *Server) Home(ctx *gin.Context) {
	Render(ctx, home.Index())
}
func (s *Server) Auth(ctx *gin.Context) {
	Render(ctx, auth.SingUp())
}
