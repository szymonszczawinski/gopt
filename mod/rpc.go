package main

import "log"

type rpc struct{}

func New() any {
	instance := new(rpc)
	return instance
}

func (s *rpc) RunService() {
	log.Println("RPC:RUN")
}
func (s *rpc) StopService() {
	log.Println("RPC:RUN")
}
