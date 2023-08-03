package http

import (
	"context"
	"embed"
	"fmt"
	"gosi/coreapi/viewcon"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

const (
	layoutsDir = "public/layouts"
)

type StaticContent struct {
	PublicDir embed.FS
}

type IHttpServer interface {
	AddController(c viewcon.IController)
}

type Routes struct {
	root  *gin.RouterGroup
	pages *gin.RouterGroup
	api   *gin.RouterGroup
}
type httpServer struct {
	server     *http.Server
	router     *gin.Engine
	routes     Routes
	renderrer  multitemplate.Renderer
	fileSystem embed.FS
	group      *errgroup.Group
	ctx        context.Context
}

func NewHttpServer(context context.Context, group *errgroup.Group, port int, staticContent StaticContent) *httpServer {
	renderrer := multitemplate.NewRenderer()
	ginRouter := createGinRouter(staticContent.PublicDir, renderrer)
	root, pages, api := configureMainRoutes(ginRouter)
	instance := httpServer{
		server: &http.Server{
			Addr:    fmt.Sprintf("localhost:%v", port),
			Handler: ginRouter,
		},
		router: ginRouter,
		routes: Routes{
			root:  root,
			pages: pages,
			api:   api,
		},
		renderrer:  renderrer,
		fileSystem: staticContent.PublicDir,
		group:      group,
		ctx:        context,
	}
	return &instance
}

func (self *httpServer) Start() {
	self.router.HTMLRender = self.renderrer
	self.group.Go(func() error {

		self.group.Go(func() error {
			// service connections
			if err := self.server.ListenAndServe(); err != nil {
				log.Printf("Listen: %s\n", err)
				return err
			}
			return nil
		})
		<-self.ctx.Done()
		ctx, cancel := context.WithTimeout(self.ctx, 5*time.Second)
		// Listen for the interrupt signal.
		defer cancel()
		if err := self.server.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
		return nil
	})
}

func (self *httpServer) AddController(c viewcon.IController) {
	c.ConfigureRoutes(self.routes.root, self.routes.pages, self.routes.api, self.fileSystem)
	self.renderrer = c.LoadViews(self.renderrer)
}
func createGinRouter(fs embed.FS, renderrer multitemplate.Renderer) *gin.Engine {
	engine := gin.Default()
	// engine.HTMLRender = loadTemplates(fs, renderrer)
	engine.StaticFS("/public", http.FS(fs))
	return engine
}
