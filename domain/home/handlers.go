package home

import (
	"embed"
	"gosi/coreapi/viewhandlers"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var (
	HomeView  = viewhandlers.View{Name: "home", Template: "public/home/home.html"}
	ErrorView = viewhandlers.View{Name: "error", Template: "public/error/error.html"}
)

type homeHandler struct {
	viewhandlers.BaseHandler
}

func NewHomeHandler(fs embed.FS) *homeHandler {
	instance := homeHandler{
		BaseHandler: viewhandlers.BaseHandler{
			FileSystem: fs,
		},
	}
	return &instance
}

func (handler *homeHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	routes.Root().GET("/", handler.homePage)
}

func (handler *homeHandler) LoadViews(r multitemplate.Renderer) multitemplate.Renderer {
	viewhandlers.AddCompositeView(r, HomeView.Name, HomeView.Template, viewhandlers.GetLayouts(), handler.FileSystem)
	viewhandlers.AddCompositeView(r, ErrorView.Name, ErrorView.Template, viewhandlers.GetLayouts(), handler.FileSystem)
	return r
}

func (handler *homeHandler) homePage(c *gin.Context) {
	c.HTML(http.StatusOK, HomeView.Name, gin.H{
		"title": "HOME",
	})
}
