package service

import (
	"errors"
	"fmt"
	"log"
	"sync"
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
func newServiceManager() *serviceManager {
	serviceManager := new(serviceManager)
	serviceManager.services = map[ServiceType]any{}
	return serviceManager
}

func GetServiceManager() *serviceManager {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			log.Println("Creating single instance now.")
			singleInstance = newServiceManager()
		} else {
			log.Println("Single instance already created.")
		}
	} else {
		log.Println("Single instance already created.")
	}

	return singleInstance
}
