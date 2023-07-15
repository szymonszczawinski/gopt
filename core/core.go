package core

import (
	"gosi/core/http"
	"gosi/core/messenger"
	"gosi/core/storage/bun"
	"gosi/issues/storage"

	"gosi/core/service"
	iservice "gosi/coreapi/service"
	projectservice "gosi/issues/service"

	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"plugin"
	"syscall"

	"golang.org/x/sync/errgroup"
)

var systemStartParameters map[string]any

func Start(cla map[string]any) {
	log.Println("START CORE")
	systemStartParameters = cla
	baseContext, cancel := context.WithCancel(context.Background())
	signalChannel := registerShutdownHook(cancel)
	mainGroup, groupContext := errgroup.WithContext(baseContext)
	service.NewServiceManager(mainGroup, groupContext)
	//some simple comment
	startServices(mainGroup, groupContext)
	// time.Sleep(time.Second * 5)
	if err := mainGroup.Wait(); err == nil {
		log.Println("FINISH CORE")
	}

	defer close(signalChannel)
}

func startServices(eg *errgroup.Group, ctx context.Context) {
	log.Println("START CORE :: START SERVICES")

	startCoreServices(eg, ctx)
}

func startCoreServices(eg *errgroup.Group, ctx context.Context) {
	log.Println("START CORE :: START CORE SERVICES")
	sm, _ := service.GetServiceManager()

	log.Println("Starting MESSENGER SERVICE")
	messengerService := messenger.NewMessengerService(eg, ctx)
	sm.StartComponent(iservice.ComponentTypeMessenger, messengerService)

	log.Println("Starting DATABASE")
	databaseConnection := bun.NewBunDatabase(eg, ctx)
	sm.StartComponent(iservice.ComponentTypeBunDatabase, databaseConnection)

	log.Println("Starting ISSUE REPOSITORY")
	issueRepository := storage.NewIssueRepository(eg, ctx, databaseConnection)
	sm.StartComponent(iservice.ComponentTypeIssueRepository, issueRepository)

	log.Println("Starting ISSUE SERVICE")
	issueService := projectservice.NewProjectService(eg, ctx, issueRepository)
	sm.StartComponent(iservice.ComponentTypeIssueService, issueService)

	// log.Println("Starting STORAGE SERVICE")
	// storageService := storage.NewStorageService(eg, ctx, repository)
	// sm.StartService(iservice.ServiceTypeIStorageService, storageService)

	log.Println("Starting HTTP SERVER SERVICE")
	httpServerService := http.NewHttpServerService(eg, ctx)
	sm.StartComponent(iservice.ComponentTypeHttpServerService, httpServerService)

	log.Println("Starting HTTP CLIENT SERVICE")
	httpClientService := http.NewHttpClientService(eg, ctx)
	sm.StartComponent(iservice.ComponentTypeHttpClientService, httpClientService)

}

func createPluginService(serviceLocation string, serviceName string) iservice.IComponent {
	plug, err := plugin.Open(serviceLocation)
	if err != nil {
		log.Println("Could not load: ", serviceName, "Error: ", err)
		return nil
	}
	createMethod, err := plug.Lookup(iservice.NEW_FUNCTION)
	if err != nil {
		log.Println("Could not get New from: ", serviceName)
		return nil
	}
	createFunction, isCreateFunction := createMethod.(func() any)
	if !isCreateFunction {
		log.Println(fmt.Sprintf("Not ceate function %T", createMethod))
		return nil
	}
	instance := createFunction()
	serviceInstance, isInstance := instance.(iservice.IComponent)
	if !isInstance {
		log.Println("Instance is not IModService")
		return nil
	}
	return serviceInstance
}

func registerShutdownHook(cancel context.CancelFunc) chan os.Signal {
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		cancel()
	}()

	return sigCh

}
