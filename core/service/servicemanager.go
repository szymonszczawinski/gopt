package service

import (
	"context"
	"errors"
	"fmt"
	"gopt/coreapi/queue"
	"gopt/coreapi/service"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	smLock         = &sync.Mutex{}
	singleInstance *serviceManager
)

type serviceManager struct {
	components map[service.ComponentType]service.IComponent
	looper     queue.IJobQueue
}

func NewServiceManager(eg *errgroup.Group, ctx context.Context) service.IServiceManager {
	smLock.Lock()
	defer smLock.Unlock()
	instance := new(serviceManager)
	instance.components = map[service.ComponentType]service.IComponent{}
	instance.looper = queue.NeqJobQueue("serviceManager", eg)
	instance.looper.Start(ctx)
	singleInstance = instance
	return instance
}

func GetServiceManager() (service.IServiceManager, error) {
	smLock.Lock()
	defer smLock.Unlock()
	if singleInstance == nil {
		return nil, errors.New("no Service Manager created")
	}
	return singleInstance, nil
}

// Get Service of given serviceType or return error if service is not registerred
func (sm *serviceManager) MustGetComponent(componentType service.ComponentType) service.IComponent {
	component, componentExists := sm.components[componentType]
	if componentExists {
		log.Println("Return component of type: ", componentType)
		return component
	}
	log.Fatalf("Componnent %v does not exists", componentType)
	return nil
}

// Register given service for a given ServiceType as a key and start service
func (sm *serviceManager) StartComponent(componentType service.ComponentType, component service.IComponent) error {
	_, componentExists := sm.components[componentType]
	if componentExists {
		return fmt.Errorf("component %v already registerred", componentType)
	}
	sm.RegisterComponent(componentType, component)
	component.StartComponent()
	return nil
}

// Register given service for a given ServiceType
func (sm *serviceManager) RegisterComponent(componentType service.ComponentType, component service.IComponent) {
	log.Println("Register Component: ", componentType)
	sm.components[componentType] = component
}
