package dummy

import (
	"context"
	"fmt"
	"gopt/coreapi/queue"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

type IDummyService interface {
	VoidMethod(message string)
}

type dummyService struct {
	looper queue.IJobQueue
}

func (s *dummyService) VoidMethod(message string) {
	log.Println("dummyService::VoidMethod")
	s.looper.Add(&queue.Job{Execute: func() {
		time.Sleep(time.Second)
		fmt.Println("Hello ", message)

	}})
}

func NewDummyService(eg *errgroup.Group, ctx context.Context) *dummyService {
	instance := new(dummyService)
	instance.looper = queue.NeqJobQueue("dummy", eg)
	instance.looper.Start(ctx)
	// instance.looper.Wait()
	return instance
}
