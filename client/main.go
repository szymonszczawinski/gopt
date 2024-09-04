package main

import (
	"context"
	"gopt/client/config"
	"gopt/client/connector/http"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	slog.Info("gopt :: CLIENT :: START")
	baseContext, cancel := context.WithCancel(context.Background())
	signalChannel := registerShutdownHook(cancel)
	mainGroup, groupContext := errgroup.WithContext(baseContext)

	configMap, err := config.GetClientConfig()
	if err != nil {
		slog.Info("Could not load configuration", err)
		return
	}
	serverPort := configMap[config.HTTP_SERVER_PORT]
	s64 := serverPort.(float64)
	httpServer := http.NewHttpServer(groupContext, mainGroup, int(s64))
	httpServer.Start()

	if err := mainGroup.Wait(); err == nil {
		slog.Info("FINISH CORE")
	}

	defer close(signalChannel)
}

func registerShutdownHook(cancel context.CancelFunc) chan os.Signal {
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-sigCh
		slog.Info("Shutdown Hook triggered")
		cancel()
	}()

	return sigCh
}
