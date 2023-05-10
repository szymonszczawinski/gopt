package service

import (
	"context"
	"coreapi/queue"
	"errors"
	"fmt"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

var singleInstance *serviceManager
var lock = &sync.Mutex{}

type ServiceType string
type IServceManager interface {
	GetService(serviceType ServiceType) (any, error)
	AddService(serviceType ServiceType, service any) error
}
type serviceManager struct {
	services map[ServiceType]any
	looper   queue.JobQueue
}

func (s *serviceManager) GetService(serviceType ServiceType) (any, error) {
	service, serviceExists := s.services[serviceType]
	if serviceExists {
		log.Println("Return service of type: ", serviceType, " :: ", service)
		return service, nil
	}
	return nil, errors.New(fmt.Sprintf("Service: %v is not registerred", serviceType))
}
func (s *serviceManager) AddService(serviceType ServiceType, service any) error {
	_, serviceExists := s.services[serviceType]
	if serviceExists {
		return errors.New(fmt.Sprintf("Service %v already registerred", serviceType))
	}
	s.services[serviceType] = service
	log.Println("Service added: ", serviceType)
	return nil
}
func NewServiceManager(eg *errgroup.Group, ctx context.Context) *serviceManager {
	instance := new(serviceManager)
	instance.services = map[ServiceType]any{}
	instance.looper = *queue.NeqJobQueue("serviceManager", eg)
	instance.looper.Start(ctx)
	singleInstance = instance
	return instance
}

func GetServiceManager() (*serviceManager, error) {
	if singleInstance == nil {
		return nil, errors.New("No Service Manager created")
	}
	return singleInstance, nil
}
