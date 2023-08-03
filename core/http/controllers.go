package http

import (
	"embed"
	// "gosi/auth"
	"gosi/coreapi/viewcon"
	"log"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type homeController struct {
	viewcon.Controller
}

func NewHomeController(fs embed.FS) *homeController {
	instance := homeController{
		Controller: viewcon.Controller{
			FileSystem: fs,
		},
	}
	return &instance
}

func (self *homeController) ConfigureRoutes(root, pages, api *gin.RouterGroup, fs embed.FS) {
	root.GET("/", self.homePage)
}
func (self *homeController) LoadViews(r multitemplate.Renderer) multitemplate.Renderer {
	viewcon.AddCompositeTemplate(r, "home", "public/home/home.html", viewcon.GetLayouts(), self.FileSystem)
	return r
}

func (self *homeController) homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{
		"title": "HOME",
	})
}

func configureMainRoutes(router *gin.Engine) (*gin.RouterGroup, *gin.RouterGroup, *gin.RouterGroup) {
	rootRoute := router.Group("/gosi")
	apiRoute := rootRoute.Group("/api")
	pagesRoute := rootRoute.Group("/pages")

	// apiRoute.Use(auth.SessionAuth())
	// pagesRoute.Use(auth.SessionAuth())
	return rootRoute, pagesRoute, apiRoute
}

func root(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		routes := router.Routes()
		routesMap := map[string]string{}

		log.Println(routes)
		for _, r := range routes {
			routesMap[r.Path] = r.Handler
		}
		log.Println(routesMap)

		c.String(http.StatusOK, "Welcome GOSI Server\nAvailable Routes:\n%v", routesMap)
	}
}
