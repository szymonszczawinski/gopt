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

var (
	smLock         = &sync.Mutex{}
	singleInstance *serviceManager
)

type serviceManager struct {
	components map[service.ComponentType]service.IComponent
	looper     queue.IJobQueue
}

// Get Service of given serviceType or return error if service is not registerred
func (self *serviceManager) MustGetComponent(componentType service.ComponentType) service.IComponent {
	component, componentExists := self.components[componentType]
	if componentExists {
		log.Println("Return component of type: ", componentType)
		return component
	}
	log.Fatal(fmt.Sprintf("Componnent %v does not exists", componentType))
	return nil
}

// Register given service for a given ServiceType as a key and start service
func (self *serviceManager) StartComponent(componentType service.ComponentType, component service.IComponent) error {
	_, componentExists := self.components[componentType]
	if componentExists {
		return errors.New(fmt.Sprintf("Component %v already registerred", componentType))
	}
	self.RegisterComponent(componentType, component)
	component.StartComponent()
	return nil
}

// Register given service for a given ServiceType
func (self *serviceManager) RegisterComponent(componentType service.ComponentType, component service.IComponent) {
	log.Println("Register Component: ", componentType)
	self.components[componentType] = component
}

func NewServiceManager(eg *errgroup.Group, ctx context.Context) *serviceManager {
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
		return nil, errors.New("No Service Manager created")
	}
	return singleInstance, nil
}
