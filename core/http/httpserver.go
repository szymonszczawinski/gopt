package http

import (
	"context"
	"embed"
	"fmt"
	common_controllers "gosi/core/http/controllers"
	issues_controllers "gosi/issues/controllers"
	user_controllers "gosi/users/controllers"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type StaticContent struct {
	PublicDir embed.FS
}

type HttpServer struct {
	server *http.Server
	engine *gin.Engine
	group  *errgroup.Group
	ctx    context.Context
}

func NewHttpServer(context context.Context, group *errgroup.Group, port int, staticContent StaticContent) *HttpServer {
	instance := new(HttpServer)
	instance.engine = gin.Default()
	instance.server = &http.Server{
		Addr:    fmt.Sprintf("localhost:%v", port),
		Handler: instance.engine,
	}
	instance.ctx = context
	instance.group = group
	// loadTemplates(instance.engine, staticContent.PublicDir)
	// instance.engine.LoadHTMLGlob("public/**/*.html")
	tmpl := template.Must(template.ParseFS(staticContent.PublicDir, "public/**/*.html"))
	for _, t := range tmpl.Templates() {
		log.Println("TEMPLATE:", t.Name())
	}
	instance.engine.SetHTMLTemplate(tmpl)
	instance.engine.StaticFS("/public", http.FS(staticContent.PublicDir))
	configureRoutes(instance.engine)
	return instance
}

func loadTemplates(engine *gin.Engine, fs embed.FS) {
	templates := multitemplate.New()
	headerTemplate, e1 := template.ParseFS(fs, "public/globals/header.tmpl")
	if e1 != nil {
		log.Println("ERROR", e1.Error())
	}
	log.Println("header loaded", headerTemplate.Name())
	ee1 := headerTemplate.Execute(os.Stdout, "")
	if ee1 != nil {
		log.Println("E:H:", ee1.Error())
	}
	templates.Add("header", headerTemplate)

	footerTemplate, e2 := template.ParseFS(fs, "public/globals/footer.tmpl")
	if e2 != nil {
		log.Println("ERROR", e2.Error())
	}
	log.Println("footer loaded", footerTemplate.Name())
	ee2 := footerTemplate.Execute(os.Stdout, "")
	if ee2 != nil {
		log.Println("E:F:", ee2.Error())
	}
	templates.Add("footer", footerTemplate)

	navTemplate, e3 := template.ParseFS(fs, "public/globals/nav.tmpl")
	if e3 != nil {
		log.Println("ERROR", e3.Error())
	}
	log.Println("nav loaded", navTemplate.Name())
	ee3 := navTemplate.Execute(os.Stdout, "")
	if ee3 != nil {
		log.Println("E:N:", ee3.Error())
	}

	templates.Add("nav", navTemplate)

	homeTemplate, e4 := template.ParseFS(fs, "public/globals/base.tmpl", "public/home/home.html")
	if e4 != nil {
		log.Println("ERROR", e4.Error())
	}
	log.Println("home loaded", homeTemplate.Name(), homeTemplate.Tree)
	ee4 := homeTemplate.Execute(os.Stdout, "")
	if ee4 != nil {
		log.Println("E:B:", ee4.Error())
	}
	templates.Add("home", homeTemplate)
	engine.HTMLRender = templates
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

func configureRoutes(router *gin.Engine) {
	rootRoute := router.Group("/gosi")
	apiRoute := rootRoute.Group("/api")
	issues_controllers.AddProjectsRoutes(apiRoute, rootRoute)
	user_controllers.AddUsersRoutes(apiRoute, rootRoute)
	common_controllers.AddBasePages(rootRoute)
}
