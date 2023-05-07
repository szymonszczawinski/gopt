package dummy

import (
	"context"
	"core/queue"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

type IService interface {
	VoidMethod(message string)
}

type dummyService struct {
	looper queue.JobQueue
}

func (s *dummyService) VoidMethod(message string) {
	log.Println("dummyService::VoidMethod")
	s.looper.Add(&queue.Job{Execute: func() { fmt.Println("Hello ", message) }})
}

func NewDummyService(eg *errgroup.Group, ctx context.Context) *dummyService {
	instance := new(dummyService)
	instance.looper = *queue.NeqJobQueue()
	instance.looper.Start(eg, ctx)
	// instance.looper.Wait()
	return instance
}
