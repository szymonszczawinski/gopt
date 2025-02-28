package messenger

import (
	"context"
	"gopt/core/service"
	"gopt/coreapi"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

const (
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

type messengerService struct {
	ctx      context.Context
	looper   coreapi.IJobQueue
	handlers map[coreapi.Topic][]IMessengerHandler
}

func NewMessengerService(eg *errgroup.Group, ctx context.Context) coreapi.IMessenger {
	slog.Info("New Messenger Service")
	messenger := new(messengerService)
	messenger.ctx = ctx
	messenger.looper = coreapi.NeqJobQueue("messengerService", eg)
	messenger.handlers = map[coreapi.Topic][]IMessengerHandler{}
	sm, _ := service.GetServiceManager()
	sm.RegisterComponent(IMMESSENGER_HANDLER_REGISTRY, messenger)
	return messenger
}

func (s *messengerService) StartComponent() {
	slog.Info("Starting", "component", coreapi.ComponentTypeMessenger)
	s.looper.Start(s.ctx)
}

func (s *messengerService) Publish(t coreapi.Topic, m coreapi.Message, listener coreapi.PublishListener) {
	slog.Info("Publish on", "topic", t)
	go func() {
		slog.Info("Publish::GO:: ", "handlers", s.handlers)
		handlers, ok := s.handlers[t]
		if ok {
			for _, handler := range handlers {
				handler.OnPublish(t, m, listener)
			}
		} else {
			slog.Info("Could not find handlers for", "topic", t)
		}
	}()
}

func (s *messengerService) Subscribe(t coreapi.Topic, listener coreapi.SubscribeListener) {
	slog.Info("Subscribe on", "topic", t)
}

func (s *messengerService) AddHandler(t coreapi.Topic, handler IMessengerHandler) {
	slog.Info("AddHandler")
	handlers, ok := s.handlers[t]
	if ok {
		handlers = append(handlers, handler)
	} else {
		s.handlers[t] = []IMessengerHandler{handler}
	}
	slog.Info("andler added for", "topic", t, " added")
}

func (s *messengerService) RemoveHandler(handler IMessengerHandler) {
	slog.Info("RemoveHandler")
}
