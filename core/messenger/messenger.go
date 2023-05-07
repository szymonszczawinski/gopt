package messenger

import (
// "fmt"
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
type IMessengerHalndlerRegistry interface {
	AddHandler(t Topic, handler IMessengerHandler)
	RemoveHandler(handler IMessengerHandler)
}
type IMessenger interface {
	Publish(t Topic, m Message, listener PublishListener)
	Subscribe(t Topic, listener SubscribeListener)
}

type Messenger struct {
	handlers map[Topic][]IMessengerHandler
}
