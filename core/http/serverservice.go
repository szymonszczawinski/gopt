package http

import (
	"context"
	"gosi/coreapi/queue"
	"gosi/coreapi/service"
	"log"

	"golang.org/x/sync/errgroup"
)

type httpServerService struct {
	looper queue.IJobQueue
	ctx    context.Context
	server HttpServer
}

func NewHttpServerService(eg *errgroup.Group, ctx context.Context) *httpServerService {
	instance := new(httpServerService)
	instance.ctx = ctx
	instance.looper = queue.NeqJobQueue("httpServerService", eg)
	instance.server = *NewHttpServer(ctx, eg, 8081)

	return instance
}

func (s *httpServerService) StartService() {
	log.Println("Starting", service.ServiceTypeIHttpServerService)
	s.looper.Start(s.ctx)
	s.server.Start()
}
