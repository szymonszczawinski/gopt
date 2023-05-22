package messenger

import (
	"context"
	"gosi/core/service"
	"gosi/coreapi"
	"gosi/coreapi/queue"
	"log"

	"golang.org/x/sync/errgroup"
)

// "fmt"

const (
	IMESSENGER                   = "IMessenger"
	IMMESSENGER_HANDLER_REGISTRY = "IMessengerHalndlerRegistry"
)

type IMessengerHandler interface {
	OnPublish(t coreapi.Topic, m coreapi.Message, listener coreapi.PublishListener)
	OnSubscribe(t coreapi.Topic, listener coreapi.SubscribeListener)
}
type IMessengerHandlerRegistry interface {
	AddHandler(t coreapi.Topic, handler IMessengerHandler)
	RemoveHandler(handler IMessengerHandler)
}
type IMessenger interface {
	Publish(t coreapi.Topic, m coreapi.Message, listener coreapi.PublishListener)
	Subscribe(t coreapi.Topic, listener coreapi.SubscribeListener)
}

type messengerService struct {
	ctx      context.Context
	looper   queue.JobQueue
	handlers map[coreapi.Topic][]IMessengerHandler
}

func NewMessengerService(eg *errgroup.Group, ctx context.Context) *messengerService {
	log.Println("New Messenger Service")
	messenger := new(messengerService)
	messenger.ctx = ctx
	messenger.looper = *queue.NeqJobQueue("messengerService", eg)
	messenger.handlers = make(map[coreapi.Topic][]IMessengerHandler)
	sm, _ := service.GetServiceManager()
	sm.RegisterService(IMMESSENGER_HANDLER_REGISTRY, messenger)
	return messenger
}

func (s *messengerService) StartService() {
	log.Println("Starting", IMESSENGER)
	s.looper.Start(s.ctx)
}

func (s *messengerService) Publish(t coreapi.Topic, m coreapi.Message, listener coreapi.PublishListener) {
	log.Println("Publish on topic:", t)
	go func() {
		log.Println("Publish::GO::handlers: ", s.handlers)
		handlers, ok := s.handlers[t]
		if ok {
			for _, handler := range handlers {
				handler.OnPublish(t, m, listener)
			}
		} else {
			log.Println("Could not find handlers for topic:", t)
		}
	}()
}

func (s *messengerService) Subscribe(t coreapi.Topic, listener coreapi.SubscribeListener) {

}
func (s *messengerService) AddHandler(t coreapi.Topic, handler IMessengerHandler) {
	log.Println("AddHandler")
	handlers, ok := s.handlers[t]
	if ok {
		handlers = append(handlers, handler)
	} else {
		s.handlers[t] = []IMessengerHandler{handler}
	}
	log.Println("Handler for topic:", t, " added")
}

func (s *messengerService) RemoveHandler(handler IMessengerHandler) {
	log.Println("RemoveHandler")
}
