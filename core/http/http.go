package http

import (
	"core/messenger"
	"core/service"
	"coreapi"
	"fmt"
	"log"
)

const ()

type httpService struct{}

func NewHttpService() *httpService {
	serviceInstance := new(httpService)
	res, err := service.GetServiceManager().GetService(messenger.IMMESSENGER_HANDLER_REGISTRY)
	if err == nil {
		impl, ok := res.(messenger.IMessengerHandlerRegistry)
		if ok {
			impl.AddHandler(coreapi.HELLO, serviceInstance)
		} else {
			log.Println("Incorrect type", impl)
		}
	} else {
		log.Println("Could not find service: ", messenger.IMMESSENGER_HANDLER_REGISTRY)
	}
	log.Println("New Http Service")
	return serviceInstance
}

func (s *httpService) OnPublish(t coreapi.Topic, m coreapi.Message, l coreapi.PublishListener) {
	log.Println(fmt.Sprintf("Message: %v published on topic: %v", m, t))
}
func (s *httpService) OnSubscribe(t coreapi.Topic, listener coreapi.SubscribeListener) {
	log.Println(fmt.Sprintf("Subscribe request on topic: %v", t))
}
