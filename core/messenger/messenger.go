package messenger

import (
	"context"
	"gopt/core/service"
	imessenger "gopt/coreapi/messenger"
	"gopt/coreapi/queue"
	iservice "gopt/coreapi/service"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

const (
	IMMESSENGER_HANDLER_REGISTRY = "IMessengerHalndlerRegistry"
)

type IMessengerHandler interface {
	OnPublish(t imessenger.Topic, m imessenger.Message, listener imessenger.PublishListener)
	OnSubscribe(t imessenger.Topic, listener imessenger.SubscribeListener)
}
type IMessengerHandlerRegistry interface {
	AddHandler(t imessenger.Topic, handler IMessengerHandler)
	RemoveHandler(handler IMessengerHandler)
}

type messengerService struct {
	ctx      context.Context
	looper   queue.IJobQueue
	handlers map[imessenger.Topic][]IMessengerHandler
}

func NewMessengerService(eg *errgroup.Group, ctx context.Context) imessenger.IMessenger {
	slog.Info("New Messenger Service")
	messenger := new(messengerService)
	messenger.ctx = ctx
	messenger.looper = queue.NeqJobQueue("messengerService", eg)
	messenger.handlers = map[imessenger.Topic][]IMessengerHandler{}
	sm, _ := service.GetServiceManager()
	sm.RegisterComponent(IMMESSENGER_HANDLER_REGISTRY, messenger)
	return messenger
}

func (s *messengerService) StartComponent() {
	slog.Info("Starting", "component", iservice.ComponentTypeMessenger)
	s.looper.Start(s.ctx)
}

func (s *messengerService) Publish(t imessenger.Topic, m imessenger.Message, listener imessenger.PublishListener) {
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

func (s *messengerService) Subscribe(t imessenger.Topic, listener imessenger.SubscribeListener) {
	slog.Info("Subscribe on", "topic", t)
}

func (s *messengerService) AddHandler(t imessenger.Topic, handler IMessengerHandler) {
	slog.Info("AddHandler")
	handlers, ok := s.handlers[t]
	if ok {
		handlers = append(handlers, handler)
	} else {
		s.handlers[t] = []IMessengerHandler{handler}
	}
	slog.Info("andler added for", "topic", t)
}

func (s *messengerService) RemoveHandler(handler IMessengerHandler) {
	slog.Info("RemoveHandler")
}
