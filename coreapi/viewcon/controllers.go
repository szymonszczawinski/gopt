package viewcon

import (
	"embed"

	"github.com/gin-gonic/gin"
)

type Action func(c *gin.Context)

type IController interface {
	Perform(a Action) gin.HandlerFunc
	ConfigureRoutes(root, pages, api *gin.RouterGroup, fs embed.FS)
}

type Controller struct {
	FileSystem embed.FS
}

func (self *Controller) Perform(a Action) gin.HandlerFunc {
	return func(c *gin.Context) {
		a(c)
	}
}
