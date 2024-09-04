package controllers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		routes := router.Routes()
		routesMap := map[string]string{}

		slog.Info("handle root", "routes", routes)
		for _, r := range routes {
			routesMap[r.Path] = r.Handler
		}
		slog.Info("handle root", "routes map", routesMap)

		c.String(http.StatusOK, "Welcome gopt Client\nAvailable Routes:\n%v", routesMap)
	}
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

func Api(c *gin.Context) {
	api := map[string]any{}
	library := map[string]any{}
	library["/gopt/api/books/showall"] = "Show All Books"
	api["library"] = library

	c.JSON(http.StatusOK, gin.H{"api": api})
}
