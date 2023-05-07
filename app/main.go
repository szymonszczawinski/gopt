package main

import (
	"context"
	"core/dummy"
	"core/http"
	"core/messenger"
	// "core/queue"
	"core/service"
	// "fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	log.Println("Hello JOSI")
	// start()
	message := "string message"
	tryJob(message)
}

func tryJob(message string) {
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		cancel()
	}()
	dummYservice := dummy.NewDummyService(ctx)
	dummYservice.VoidMethod("szymon")
	// jq := queue.NeqJobQueue()
	// jq.Start(ctx)
	// j := queue.Job{Execute: func() { fmt.Println(message) }}
	// jq.Add(&j)
	// jq.Wait()
}

func start() {
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {

		serviceManager := service.GetServiceManager()
		messengerService := messenger.NewMessenger()
		serviceManager.AddService(messenger.IMESSENGER, messengerService)
		serviceManager.AddService(messenger.IMMESSENGER_HANDLER_REGISTRY, messengerService)
		http.NewHttpService()
		messengerService.Publish(messenger.HELLO, "Szymon", nil)
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatal("Error", err)
	}
	log.Println("Completed successfully!")
}
