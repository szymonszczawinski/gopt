package dummy

import (
	"context"
	"fmt"
	"gopt/coreapi"
	"log/slog"
	"time"

	"golang.org/x/sync/errgroup"
)

type IDummyService interface {
	VoidMethod(message string)
}

type dummyService struct {
	looper coreapi.IJobQueue
}

func (s *dummyService) VoidMethod(message string) {
	slog.Info("dummyService::VoidMethod")
	s.looper.Add(&coreapi.Job{Execute: func() {
		time.Sleep(time.Second)
		fmt.Println("Hello ", message)
	}})
}

func NewDummyService(eg *errgroup.Group, ctx context.Context) *dummyService {
	instance := new(dummyService)
	instance.looper = coreapi.NeqJobQueue("dummy", eg)
	instance.looper.Start(ctx)
	// instance.looper.Wait()
	return instance
}
