package http

import (
	"embed"
	// "gosi/auth"
	"gosi/coreapi/viewhandlers"
	"log"
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

func configureMainRoutes(router *gin.Engine) *viewhandlers.Routes {
	rootRoute := router.Group("/gosi")
	apiRoute := rootRoute.Group("/api")
	viewsRoute := rootRoute.Group("/views")

	// apiRoute.Use(auth.SessionAuth)
	// viewsRoute.Use(auth.SessionAuth)
	routes := viewhandlers.NewRoutes(rootRoute, viewsRoute, apiRoute)
	return routes
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
