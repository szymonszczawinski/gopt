package viewhandlers

import (
	"embed"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	root  *gin.RouterGroup
	views *gin.RouterGroup
	apis  *gin.RouterGroup
}

func NewRoutes(root, views, apis *gin.RouterGroup) *Routes {
	return &Routes{root: root,
		views: views,
		apis:  apis}
}

func (self Routes) Root() *gin.RouterGroup {
	return self.root
}

func (self Routes) Views() *gin.RouterGroup {
	return self.views
}

func (self Routes) Apis() *gin.RouterGroup {
	return self.apis
}

type Action func(c *gin.Context)

type IViewHandler interface {
	Perform(a Action) gin.HandlerFunc
	ConfigureRoutes(routes Routes)
	LoadViews(r multitemplate.Renderer) multitemplate.Renderer
}

type ApiHandler interface {
	ConfigureApiRoutes(routes *Routes)
}

type BaseHandler struct {
	FileSystem embed.FS
}

func (self *BaseHandler) Perform(a Action) gin.HandlerFunc {
	return func(c *gin.Context) {
		a(c)
	}
}
