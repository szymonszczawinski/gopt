package connector

import ()

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

type IConnector interface {
	Open(callback IConnectorCallback)
}
type IConnectorSender interface {
	SendMessage(where any, message any) any
}
type IConnectorCallback interface {
	OnMessage(from any, message any)
}

type IHttpConnector interface {
	PerformRequest(method HttpMethod, url string, payload any, callback IConnectorCallback)
}
