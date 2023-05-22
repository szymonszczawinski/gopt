package http

import (
	"context"
	"fmt"
	"gosi/core/connector/http/controllers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type HttpServer struct {
	server *http.Server
	engine *gin.Engine
	group  *errgroup.Group
	ctx    context.Context
}

func NewHttpServer(context context.Context, group *errgroup.Group, port int) *HttpServer {
	instance := new(HttpServer)
	instance.engine = gin.Default()
	instance.server = &http.Server{
		Addr:    fmt.Sprintf("localhost:%v", port),
		Handler: instance.engine,
	}
	instance.ctx = context
	instance.group = group
	configureRoutes(instance.engine)
	return instance
}
func configureRoutes(router *gin.Engine) {
	router.GET("/gosi", controllers.Root(router))
	router.GET("/gosi/hello", controllers.Hello)

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
