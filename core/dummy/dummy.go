package dummy

import (
	"context"
	"core/queue"
	"fmt"
	"log"
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

func NewDummyService(ctx context.Context) *dummyService {
	instance := new(dummyService)
	instance.looper = *queue.NeqJobQueue()
	instance.looper.Start(ctx)
	// instance.looper.Wait()
	return instance
}
