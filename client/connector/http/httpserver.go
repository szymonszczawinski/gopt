package http

import (
	"context"
	"fmt"
	"gopt/client/connector/http/controllers"
	"log"
	"log/slog"
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
	port   int
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
	instance.port = port
	configureRoutes(instance.engine)
	return instance
}

func configureRoutes(router *gin.Engine) {
	router.GET("/gopt", controllers.Root(router))
	router.GET("/gopt/hello", controllers.Hello)
	router.GET("/gopt/api", controllers.Api)
}

func (s *HttpServer) Start() {
	s.group.Go(func() error {
		// service connections
		slog.Info("Sever started on port:", s.port)
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
	slog.Info("Server exiting")
}
