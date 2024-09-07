package viewhandlers

import (
	"github.com/gin-gonic/gin"
)

type Routes struct {
	root  *gin.RouterGroup
	views *gin.RouterGroup
	apis  *gin.RouterGroup
}

func NewRoutes(root, views, apis *gin.RouterGroup) *Routes {
	return &Routes{
		root:  root,
		views: views,
		apis:  apis,
	}
}

func (r Routes) Root() *gin.RouterGroup {
	return r.root
}

func (r Routes) Views() *gin.RouterGroup {
	return r.views
}

func (r Routes) Apis() *gin.RouterGroup {
	return r.apis
}

// type Action func(c *gin.Context)

type IViewHandler interface {
	// Perform(a Action) gin.HandlerFunc
	ConfigureRoutes(routes Routes)
}

// type ApiHandler interface {
// ConfigureApiRoutes(routes *Routes)
// }

// type BaseHandler struct {
// FileSystem embed.FS
// }

// func (h *BaseHandler) Perform(a Action) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		a(c)
// 	}
// }
