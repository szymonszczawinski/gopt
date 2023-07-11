package http

import (
	"context"
	"fmt"
	"gosi/core/messenger"
	"gosi/core/service"
	imessenger "gosi/coreapi/messenger"
	"gosi/coreapi/queue"
	"log"

	"golang.org/x/sync/errgroup"
)

const (
	IHTTP_CLIENT_SERVICE = "IHttpClientService"
)

type httpClientService struct {
	looper queue.IJobQueue
	ctx    context.Context
}

func NewHttpClientService(eg *errgroup.Group, ctx context.Context) *httpClientService {
	serviceInstance := new(httpClientService)
	serviceInstance.ctx = ctx
	serviceInstance.looper = queue.NeqJobQueue("httpClientService", eg)
	sm, err := service.GetServiceManager()
	if err == nil {

		res, err := sm.GetService(messenger.IMMESSENGER_HANDLER_REGISTRY)
		if err == nil {
			impl, ok := res.(messenger.IMessengerHandlerRegistry)
			if ok {
				impl.AddHandler(imessenger.HELLO, serviceInstance)
			} else {
				log.Println("Incorrect type", impl)
			}
		} else {
			log.Println("Could not find service: ", messenger.IMMESSENGER_HANDLER_REGISTRY)
		}
	}
	return serviceInstance
}

func (s *httpClientService) StartService() {
	log.Println("Starting", IHTTP_CLIENT_SERVICE)
	s.looper.Start(s.ctx)
}

func (s *httpClientService) OnPublish(t imessenger.Topic, m imessenger.Message, l imessenger.PublishListener) {
	log.Println(fmt.Sprintf("Message: %v published on topic: %v", m, t))
}
func (s *httpClientService) OnSubscribe(t imessenger.Topic, listener imessenger.SubscribeListener) {
	log.Println(fmt.Sprintf("Subscribe request on topic: %v", t))
}
