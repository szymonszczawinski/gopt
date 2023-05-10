package main

import (
	"coreapi"
	"log"
)

type rpc struct {
	broker *coreapi.MessageBroker
}

func New() any {
	instance := new(rpc)
	instance.broker = coreapi.NewBroker()
	return instance
}

func (s *rpc) StartService() {
	log.Println("RPC:RUN")
	s.broker.Publish(coreapi.HELLO, coreapi.Message("Hello"), nil)
}
func (s *rpc) StopService() {
	log.Println("RPC:RUN")
}
