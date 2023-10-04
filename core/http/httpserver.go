package http

import (
	"context"
	"embed"
	"fmt"
	"gosi/coreapi/viewhandlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	AddHandler(c viewhandlers.IViewHandler)
}

type httpServer struct {
	server     *http.Server
	router     *gin.Engine
	routes     *viewhandlers.Routes
	renderrer  multitemplate.Renderer
	fileSystem embed.FS
	group      *errgroup.Group
	ctx        context.Context
}

func NewHttpServer(context context.Context, group *errgroup.Group, port int, staticContent StaticContent) *httpServer {
	renderrer := multitemplate.NewRenderer()
	ginRouter := createGinRouter(staticContent.PublicDir, renderrer)
	routes := configureMainRoutes(ginRouter)
	instance := httpServer{
		server: &http.Server{
			Addr:    fmt.Sprintf("localhost:%v", port),
			Handler: ginRouter,
		},
		router:     ginRouter,
		routes:     routes,
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

func (self *httpServer) AddHandler(c viewhandlers.IViewHandler) {
	c.ConfigureRoutes(*self.routes)
	self.renderrer = c.LoadViews(self.renderrer)
}
func createGinRouter(fs embed.FS, renderrer multitemplate.Renderer) *gin.Engine {
	engine := gin.Default()
	// engine.HTMLRender = loadTemplates(fs, renderrer)
	engine.StaticFS("/public", http.FS(fs))
	cookieOptions := sessions.Options{Path: "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Domain:   "",
		MaxAge:   60 * 5,
	}
	cookieStore := cookie.NewStore([]byte(os.Getenv("SECRET")))
	cookieStore.Options(cookieOptions)
	engine.Use(sessions.Sessions("mysession", cookieStore))
	return engine
}
