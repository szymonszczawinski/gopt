package service

import (
	"errors"
	"fmt"
)

type ServiceType string
type IServceManager interface {
	GetService(serviceType ServiceType) (any, error)
	AddService(serviceType ServiceType, service any)
}
type ServiceManager struct {
	services map[ServiceType]any
}

func (s *ServiceManager) GetService(serviceType ServiceType) (any, error) {
	service, serviceExists := s.services[serviceType]
	if serviceExists {
		return service, nil
	}
	return nil, errors.New(fmt.Sprintf("Service: %v is not registerred", serviceType))
}
