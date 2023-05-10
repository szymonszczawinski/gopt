package core

import (
	"core/config"
	"core/dummy"
	"core/service"

	"rpc"

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
	dummy.NewDummyService(mainGroup, groupContext)

	// startServices(mainGroup, groupContext)
	// time.Sleep(time.Second * 5)
	if err := mainGroup.Wait(); err == nil {
		log.Println("FINISH CORE")
	}

	defer close(signalChannel)
}

func startServices(eg *errgroup.Group, ctx context.Context) {

	startCoreServices(eg, ctx)
	// startModServices(eg, ctx)
}
func startCoreServices(eg *errgroup.Group, ctx context.Context) {

}
func startModServices(eg *errgroup.Group, ctx context.Context) {
	service := createModService("RPC", "../mod/rpc/rpc.so", eg, ctx)
	if service != nil {
		service.StartService()
	}

}
func createModService(serviceName string, serviceLocation string, eg *errgroup.Group, ctx context.Context) service.IService {
	if systemStartParameters[config.RUN_MODE] == config.RUN_MODE_PLUG {
		return createPluginService(serviceLocation, serviceName)
	} else {
		if serviceName == "RPC" {
			instance := rpc.New(eg, ctx)
			serviceInstance, isInstance := instance.(service.IService)
			if !isInstance {
				log.Println("Instance is not IModService")
				return nil
			}
			return serviceInstance
		}
	}
	return nil

}

func createPluginService(serviceLocation string, serviceName string) service.IService {
	plug, err := plugin.Open(serviceLocation)
	if err != nil {
		log.Println("Could not load: ", serviceName, "Error: ", err)
		return nil
	}
	createMethod, err := plug.Lookup(service.NEW_FUNCTION)
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
	serviceInstance, isInstance := instance.(service.IService)
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

func tryJob(message string) {
	baseContext, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		cancel()
	}()
	eg, ctx := errgroup.WithContext(baseContext)

	dummy.NewDummyService(eg, ctx)
	// dummYservice.VoidMethod("szymon")
	// dummYservice.VoidMethod("szymon")
	// dummYservice.VoidMethod("szymon")
	// dummYservice.VoidMethod("szymon")
	if err := eg.Wait(); err == nil {
		log.Println("Successfully fetched all URLs.")
	}
	// jq := queue.NeqJobQueue()
	// jq.Start(ctx)
	// j := queue.Job{Execute: func() { fmt.Println(message) }}
	// jq.Add(&j)
	// jq.Wait()
}

func start2() {
	// eg, _ := errgroup.WithContext(context.Background())
	// eg.Go(func() error {
	//
	// 	serviceManager := service.GetServiceManager()
	// 	messengerService := messenger.NewMessenger()
	// 	serviceManager.AddService(messenger.IMESSENGER, messengerService)
	// 	serviceManager.AddService(messenger.IMMESSENGER_HANDLER_REGISTRY, messengerService)
	// 	http.NewHttpService()
	// 	messengerService.Publish(coreapi.HELLO, "Szymon", nil)
	// 	return nil
	// })
	// if err := eg.Wait(); err != nil {
	// 	log.Fatal("Error", err)
	// }
	// log.Println("Completed successfully!")
}
