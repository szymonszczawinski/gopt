package main

import (
	"context"
	"gosi/client/config"
	"gosi/client/connector/http/"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("GOSI :: CLIENT :: START")
	baseContext, cancel := context.WithCancel(context.Background())
	signalChannel := registerShutdownHook(cancel)
	mainGroup, groupContext := errgroup.WithContext(baseContext)

	configMap, err := config.GetClientConfig()
	if err != nil {
		log.Println("Could not load configuration", err)
		return
	}
	serverPort := configMap[config.HTTP_SERVER_PORT]
	s64 := serverPort.(float64)
	httpServer := http.NewHttpServer(groupContext, mainGroup, int(s64))
	httpServer.Start()

	if err := mainGroup.Wait(); err == nil {
		log.Println("FINISH CORE")
	}

	defer close(signalChannel)

}

func registerShutdownHook(cancel context.CancelFunc) chan os.Signal {
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		log.Println("Shutdown Hook triggered")
		cancel()
	}()

	return sigCh

}
