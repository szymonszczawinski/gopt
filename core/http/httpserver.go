package http

import (
	"context"
	"embed"
	"fmt"
	"gosi/auth"
	common_controllers "gosi/core/http/controllers"
	issues_controllers "gosi/issues/controllers"
	user_controllers "gosi/users/controllers"
	"html/template"
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

type HttpServer struct {
	server *http.Server
	group  *errgroup.Group
	ctx    context.Context
}

func NewHttpServer(context context.Context, group *errgroup.Group, port int, staticContent StaticContent) *HttpServer {
	instance := HttpServer{
		server: &http.Server{
			Addr:    fmt.Sprintf("localhost:%v", port),
			Handler: createGinHandler(staticContent),
		},
		group: group,
		ctx:   context,
	}
	return &instance
}
func (s *HttpServer) Start() {
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

func createGinHandler(staticContent StaticContent) *gin.Engine {
	engine := gin.Default()
	engine.HTMLRender = loadTemplates(staticContent.PublicDir)
	engine.StaticFS("/public", http.FS(staticContent.PublicDir))
	configureRoutes(engine, staticContent.PublicDir)
	return engine
}

func configureRoutes(router *gin.Engine, fs embed.FS) {
	rootRoute := router.Group("/gosi")

	apiRoute := rootRoute.Group("/api")
	issues_controllers.AddProjectsRoutes(apiRoute, rootRoute)
	user_controllers.AddUsersRoutes(apiRoute, rootRoute)
	common_controllers.AddBasePages(rootRoute)

	auth.AddAuthRoutes(rootRoute, fs)

	restrictedRoute := rootRoute.Group("/restricted")
	restrictedRoute.Use(auth.SessionAuth())
	auth.AddRestrictedRoute(restrictedRoute)
}

func loadTemplates(fs embed.FS) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts := getLayouts(fs)
	addCompositeTemplate(r, "home", "public/home/home.html", layouts, fs)
	addCompositeTemplate(r, "projects", "public/projects/projects.html", layouts, fs)
	addCompositeTemplate(r, "admin", "public/admin/admin.html", layouts, fs)
	addCompositeTemplate(r, "login", "public/auth/login.html", getSimpleLayouts(), fs)
	addCompositeTemplate(r, "error", "public/error/error.html", getSimpleLayouts(), fs)
	return r
}

func addCompositeTemplate(r multitemplate.Renderer, name string, path string, layouts []string, fs embed.FS) multitemplate.Renderer {
	layouts = append(layouts, path)
	tmpl, _ := template.ParseFS(fs, layouts...)
	r.Add(name, tmpl)
	return r
}

func addSimpleTemplate(r multitemplate.Renderer, name string, path string, fs embed.FS) multitemplate.Renderer {
	tmpl, _ := template.ParseFS(fs, path)
	r.Add(name, tmpl)
	return r
}

func getLayouts(fs embed.FS) []string {
	layouts := []string{}
	site, err := embed.FS.ReadDir(fs, layoutsDir)
	if err != nil {
		panic(err.Error())
	}
	for _, layout := range site {
		layouts = append(layouts, layoutsDir+"/"+layout.Name())
	}
	return layouts
}
func getSimpleLayouts() []string {
	layouts := []string{"public/layouts/basesimple.html",
		"public/layouts/header.html",
		"public/layouts/footer.html"}
	return layouts
}
