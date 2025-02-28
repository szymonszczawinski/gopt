package http

import (
	"context"
	"gopt/core/messenger"
	"gopt/core/service"
	"gopt/coreapi"
	"log"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

const (
	IHTTP_CLIENT_SERVICE = "IHttpClientService"
)

type httpClientService struct {
	looper coreapi.IJobQueue
	ctx    context.Context
}

func NewHttpClientService(eg *errgroup.Group, ctx context.Context) *httpClientService {
	serviceInstance := new(httpClientService)
	serviceInstance.ctx = ctx
	serviceInstance.looper = coreapi.NeqJobQueue("httpClientService", eg)
	sm, err := service.GetServiceManager()
	if err == nil {

		res := sm.MustGetComponent(messenger.IMMESSENGER_HANDLER_REGISTRY)
		if err == nil {
			impl, ok := res.(messenger.IMessengerHandlerRegistry)
			if ok {
				impl.AddHandler(coreapi.HELLO, serviceInstance)
			} else {
				slog.Info("Incorrect type", impl)
			}
		} else {
			slog.Info("Could not find service: ", messenger.IMMESSENGER_HANDLER_REGISTRY)
		}
	}
	return serviceInstance
}

func (s *httpClientService) StartComponent() {
	slog.Info("Starting", IHTTP_CLIENT_SERVICE)
	s.looper.Start(s.ctx)
}

func (s *httpClientService) OnPublish(t coreapi.Topic, m coreapi.Message, l coreapi.PublishListener) {
	log.Printf("Message: %v published on topic: %v\n", m, t)
}

func (s *httpClientService) OnSubscribe(t coreapi.Topic, listener coreapi.SubscribeListener) {
	log.Printf("Subscribe request on topic: %v\n", t)
}
