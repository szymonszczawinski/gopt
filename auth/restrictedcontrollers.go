package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRestrictedRoute(route *gin.RouterGroup) {
	route.GET("admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin", gin.H{
			"title": "ADMIN",
		})
	})

}
