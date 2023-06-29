package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(router *gin.Engine) gin.HandlerFunc {
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

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

func AddBasePages(rootRoute *gin.RouterGroup) {
	rootRoute.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "HOME",
		})
	})

}
