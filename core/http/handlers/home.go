package handlers

import (
	errors "gopt/public/error"
	"gopt/public/home"

	"github.com/gin-gonic/gin"
)

type homeHandler struct{}

func NewHomeHandler() *homeHandler {
	instance := homeHandler{}
	return &instance
}

func (handler *homeHandler) ConfigureRoutes(routes Routes) {
	routes.Root().GET("/", handler.homePage)
	routes.Root().GET("/error", handler.errorPage)
}

func (handler *homeHandler) homePage(c *gin.Context) {
	isHxRequest := c.GetHeader("HX-Request")
	if isHxRequest == "true" {
		home.Home(true).Render(c.Request.Context(), c.Writer)
	} else {
		home.Home(false).Render(c.Request.Context(), c.Writer)
	}
}

func (handler *homeHandler) errorPage(c *gin.Context) {
	errors.Error("SUPER ERROR").Render(c.Request.Context(), c.Writer)
}
