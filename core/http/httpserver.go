package http

import (
	"context"
	"embed"
	"fmt"
	"gosi/coreapi/viewcon"
	"log"
	"net/http"
	"time"

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
	fileSystem embed.FS
	group      *errgroup.Group
	ctx        context.Context
}

func NewHttpServer(context context.Context, group *errgroup.Group, port int, staticContent StaticContent) *httpServer {
	ginRouter := createGinRouter(staticContent.PublicDir)
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
		fileSystem: staticContent.PublicDir,
		group:      group,
		ctx:        context,
	}
	return &instance
}

func (s *httpServer) Start() {
	s.group.Go(func() error {

		s.group.Go(func() error {
			// service connections
			if err := s.server.ListenAndServe(); err != nil {
				log.Printf("Listen: %s\n", err)
				return err
			}
			return nil
		})
		<-s.ctx.Done()
		ctx, cancel := context.WithTimeout(s.ctx, 5*time.Second)
		// Listen for the interrupt signal.
		defer cancel()
		if err := s.server.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
		return nil
	})
}

func (self *httpServer) AddController(c viewcon.IController) {
	c.ConfigureRoutes(self.routes.root, self.routes.pages, self.routes.api, self.fileSystem)
}
func createGinRouter(fs embed.FS) *gin.Engine {
	engine := gin.Default()
	engine.HTMLRender = loadTemplates(fs)
	engine.StaticFS("/public", http.FS(fs))
	return engine
}
