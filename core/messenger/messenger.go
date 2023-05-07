package messenger

import "log"

// "fmt"

const (
	IMESSENGER                   = "IMessenger"
	IMMESSENGER_HANDLER_REGISTRY = "IMessengerHalndlerRegistry"
)

type Message any
type Feedback any

type PublishListener interface {
	OnSuccess(t Topic, f Feedback)
	OnFailure(t Topic, f Feedback)
}

type SubscribeListener interface {
	OnMessage(t Topic, m Message)
}
type IMessengerHandler interface {
	OnPublish(t Topic, m Message, listener PublishListener)
	OnSubscribe(t Topic, listener SubscribeListener)
}
type IMessengerHandlerRegistry interface {
	AddHandler(t Topic, handler IMessengerHandler)
	RemoveHandler(handler IMessengerHandler)
}
type IMessenger interface {
	Publish(t Topic, m Message, listener PublishListener)
	Subscribe(t Topic, listener SubscribeListener)
}

type messengerService struct {
	handlers map[Topic][]IMessengerHandler
}

func NewMessenger() *messengerService {
	log.Println("New Messenger Service")
	messenger := new(messengerService)
	messenger.handlers = make(map[Topic][]IMessengerHandler)
	return messenger
}

func (s *messengerService) Publish(t Topic, m Message, listener PublishListener) {
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

func (s *messengerService) Subscribe(t Topic, listener SubscribeListener) {

}
func (s *messengerService) AddHandler(t Topic, handler IMessengerHandler) {
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
