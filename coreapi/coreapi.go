package coreapi

import "log"

type PublishListener interface {
	OnSuccess(t Topic, f Feedback)
	OnFailure(t Topic, f Feedback)
}

type SubscribeListener interface {
	OnMessage(t Topic, m Message)
}

type MessageBroker struct {
}

func NewBroker() *MessageBroker {
	broker := new(MessageBroker)
	return broker
}
func (broker MessageBroker) Publish(t Topic, m Message, listener PublishListener) {
	log.Println("Publish on ", t)
}
func (broker MessageBroker) Subscribe(t Topic, listener SubscribeListener) {
	log.Println("Subscribe on ", t)
}
