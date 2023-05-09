package core

import (
	"context"
	"core/dummy"
	"core/http"
	"core/messenger"
	"coreapi"
	"fmt"
	"plugin"
	"time"

	// "time"

	// "core/queue"
	"core/service"
	// "fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func Start() {
	log.Println("START CORE")

	baseContext, cancel := context.WithCancel(context.Background())
	registerShutdownHook(cancel)
	mainGroup, _ := errgroup.WithContext(baseContext)
	startServices()
	time.Sleep(time.Second * 5)
	if err := mainGroup.Wait(); err == nil {
		log.Println("FINISH CORE")
	}

}

func startServices() {
	startModSerice("RPC", "../mod/rpc/rpc.so")
}

func startModSerice(serviceName string, serviceLocation string) {
	plug, err := plugin.Open(serviceLocation)
	if err != nil {
		log.Println("Could not load: ", serviceName)
	} else {
		createMethod, err := plug.Lookup("New")
		if err != nil {
			log.Println("Could not get New from: ", serviceName)
		} else {
			createFunction, isCreateFunction := createMethod.(func() any)
			if !isCreateFunction {
				log.Println(fmt.Sprintf("Not ceate function %T", createMethod))
			} else {
				instance := createFunction()
				serviceInstance, isInstance := instance.(service.IService)
				if !isInstance {
					log.Println("Instance is not IModService")
				} else {
					serviceInstance.RunService()
				}
			}
		}
	}

}

func registerShutdownHook(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		cancel()
	}()

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

	dummYservice := dummy.NewDummyService(eg, ctx)
	dummYservice.VoidMethod("szymon")
	dummYservice.VoidMethod("szymon")
	dummYservice.VoidMethod("szymon")
	dummYservice.VoidMethod("szymon")
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
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {

		serviceManager := service.GetServiceManager()
		messengerService := messenger.NewMessenger()
		serviceManager.AddService(messenger.IMESSENGER, messengerService)
		serviceManager.AddService(messenger.IMMESSENGER_HANDLER_REGISTRY, messengerService)
		http.NewHttpService()
		messengerService.Publish(coreapi.HELLO, "Szymon", nil)
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatal("Error", err)
	}
	log.Println("Completed successfully!")
}
