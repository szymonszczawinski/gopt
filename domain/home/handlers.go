package home

import (
	"gosi/coreapi/viewhandlers"
	errors "gosi/public/error"
	"gosi/public/home"

	"github.com/gin-gonic/gin"
)

type homeHandler struct{}

func NewHomeHandler() *homeHandler {
	instance := homeHandler{}
	return &instance
}

func (handler *homeHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	routes.Root().GET("/", handler.homePage)
	routes.Root().GET("/error", handler.errorPage)
}

func (handler *homeHandler) homePage(c *gin.Context) {
	home.Home().Render(c.Request.Context(), c.Writer)
}

func (handler *homeHandler) errorPage(c *gin.Context) {
	errors.Error("SUPER ERROR").Render(c.Request.Context(), c.Writer)
}
