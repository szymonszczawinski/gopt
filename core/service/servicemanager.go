package service

import (
	"context"
	"errors"
	"fmt"
	"gopt/coreapi"
	"log"
	"log/slog"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	smLock         = &sync.Mutex{}
	singleInstance *serviceManager
)

type serviceManager struct {
	components map[coreapi.ComponentType]coreapi.IComponent
	looper     coreapi.IJobQueue
}

func NewServiceManager(eg *errgroup.Group, ctx context.Context) coreapi.IServiceManager {
	smLock.Lock()
	defer smLock.Unlock()
	instance := new(serviceManager)
	instance.components = map[coreapi.ComponentType]coreapi.IComponent{}
	instance.looper = coreapi.NeqJobQueue("serviceManager", eg)
	instance.looper.Start(ctx)
	singleInstance = instance
	return instance
}

func GetServiceManager() (coreapi.IServiceManager, error) {
	smLock.Lock()
	defer smLock.Unlock()
	if singleInstance == nil {
		return nil, errors.New("no Service Manager created")
	}
	return singleInstance, nil
}

// Get Service of given serviceType or return error if service is not registerred
func (sm *serviceManager) MustGetComponent(componentType coreapi.ComponentType) coreapi.IComponent {
	component, componentExists := sm.components[componentType]
	if componentExists {
		slog.Info("Return component", "type", componentType)
		return component
	}
	log.Fatalf("Componnent %v does not exists", componentType)
	return nil
}

// Register given service for a given ServiceType as a key and start service
func (sm *serviceManager) StartComponent(componentType coreapi.ComponentType, component coreapi.IComponent) error {
	_, componentExists := sm.components[componentType]
	if componentExists {
		return fmt.Errorf("component %v already registerred", componentType)
	}
	sm.RegisterComponent(componentType, component)
	component.StartComponent()
	return nil
}

// Register given service for a given ServiceType
func (sm *serviceManager) RegisterComponent(componentType coreapi.ComponentType, component coreapi.IComponent) {
	slog.Info("Register Component", "type", componentType)
	sm.components[componentType] = component
}
