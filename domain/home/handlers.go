package home

import (
	"embed"
	"gosi/coreapi/viewhandlers"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
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

func (self *homeHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	routes.Root().GET("/", self.homePage)
}

func (self *homeHandler) LoadViews(r multitemplate.Renderer) multitemplate.Renderer {
	viewhandlers.AddCompositeView(r, "home", "public/home/home.html", viewhandlers.GetLayouts(), self.FileSystem)
	viewhandlers.AddCompositeView(r, "error", "public/error/error.html", viewhandlers.GetLayouts(), self.FileSystem)
	return r
}

func (self *homeHandler) homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{
		"title": "HOME",
	})
}
