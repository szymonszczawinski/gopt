package http

import (
	"context"
	"gopt/coreapi/queue"
	"gopt/coreapi/service"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

type httpServerService struct {
	looper queue.IJobQueue
	ctx    context.Context
	server httpServer
}

func NewHttpServerService(eg *errgroup.Group, ctx context.Context, staticContent StaticContent) *httpServerService {
	instance := new(httpServerService)
	instance.ctx = ctx
	instance.looper = queue.NeqJobQueue("httpServerService", eg)
	instance.server = *NewHttpServer(ctx, eg, 8081, staticContent)

	return instance
}

func (s *httpServerService) StartComponent() {
	slog.Info("Starting", "component", service.ComponentTypeHttpServerService)
	s.looper.Start(s.ctx)
	s.server.Start()
}
