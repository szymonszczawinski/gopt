package rpc

import (
	"context"
	"gosi/coreapi"
	// "coreapi/queue"
	"log"

	"golang.org/x/sync/errgroup"
)

type rpc struct {
	broker *coreapi.MessageBroker
	// looper queue.JobQueue
	ctx context.Context
}

func NewRpcService(eg *errgroup.Group, ctx context.Context) any {
	instance := new(rpc)
	instance.broker = coreapi.NewBroker()
	// instance.looper = *queue.NeqJobQueue(eg)
	instance.ctx = ctx
	return instance
}

func (s *rpc) StartService() {
	log.Println("RPC:RUN")
	// s.looper.Start(s.ctx)
	s.broker.Publish(coreapi.HELLO, coreapi.Message("Hello"), nil)
}
func (s *rpc) StopService() {
	log.Println("RPC:RUN")
}
