package http

import (
	"gosi/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func configureMainRoutes(router *gin.Engine) (*gin.RouterGroup, *gin.RouterGroup, *gin.RouterGroup) {
	rootRoute := router.Group("/gosi")
	apiRoute := rootRoute.Group("/api")
	pagesRoute := rootRoute.Group("/pages")

	apiRoute.Use(auth.SessionAuth())
	pagesRoute.Use(auth.SessionAuth())
	addBasePages(rootRoute)
	return rootRoute, pagesRoute, apiRoute
	// issues_controllers.AddProjectsRoutes(apiRoute, pagesRoute)
	//
	// user_controllers.AddUsersRoutes(apiRoute, pagesRoute)
	// addBasePages(rootRoute)
	// auth.AddAuthRoutes(rootRoute, fs)

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

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

func addBasePages(rootRoute *gin.RouterGroup) {
	rootRoute.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{
			"title": "HOME",
		})
	})

}
