package service

import (
	"context"
	"errors"
	"fmt"
	"gosi/coreapi/queue"
	"gosi/coreapi/service"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

var singleInstance *serviceManager
var lock = &sync.Mutex{}

type serviceManager struct {
	services map[service.ServiceType]any
	looper   queue.IJobQueue
}

// Get Service of given serviceType or return error if service is not registerred
func (s *serviceManager) GetService(serviceType service.ServiceType) (any, error) {
	service, serviceExists := s.services[serviceType]
	if serviceExists {
		log.Println("Return service of type: ", serviceType)
		return service, nil
	}
	return nil, errors.New(fmt.Sprintf("Service: %v is not registerred", serviceType))
}

// Register given service for a given ServiceType as a key and start service
func (s *serviceManager) StartService(serviceType service.ServiceType, service service.IService) error {
	_, serviceExists := s.services[serviceType]
	if serviceExists {
		return errors.New(fmt.Sprintf("Service %v already registerred", serviceType))
	}
	s.RegisterService(serviceType, service)
	service.StartService()
	return nil
}

// Register given service for a given ServiceType
func (s *serviceManager) RegisterService(serviceType service.ServiceType, service service.IService) {
	log.Println("Register Service: ", serviceType)
	s.services[serviceType] = service
}

func NewServiceManager(eg *errgroup.Group, ctx context.Context) *serviceManager {
	instance := new(serviceManager)
	instance.services = map[service.ServiceType]any{}
	instance.looper = queue.NeqJobQueue("serviceManager", eg)
	instance.looper.Start(ctx)
	singleInstance = instance
	return instance
}

func GetServiceManager() (service.IServiceManager, error) {
	if singleInstance == nil {
		return nil, errors.New("No Service Manager created")
	}
	return singleInstance, nil
}
