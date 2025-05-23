package coreapi

type (
	Message  any
	Feedback any
)

type IMessenger interface {
	IComponent
	Publish(t Topic, m Message, listener PublishListener)
	Subscribe(t Topic, listener SubscribeListener)
}

type PublishListener interface {
	OnSuccess(t Topic, f Feedback)
	OnFailure(t Topic, f Feedback)
}

type SubscribeListener interface {
	OnMessage(t Topic, m Message)
}

var HELLO Topic = Topic{"hello"}

type Topic struct {
	name string
}
