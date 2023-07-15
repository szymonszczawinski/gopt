package messenger

import "gosi/coreapi/service"

type Message any
type Feedback any

type IMessenger interface {
	service.IComponent
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
