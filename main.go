package main

import (
	"context"
	"gosi/core"
	"gosi/core/config"
	"gosi/coreapi/queue"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("GOSI :: START")
	cla := parseCLA(os.Args)
	core.Start(cla)

	log.Println("GOSI :: FINISH")
	// sqlite.GetSqliteRepository()
	// RunTests()
}

func parseCLA(args []string) map[string]any {
	cla := map[string]any{}
	cla[config.RUN_MODE] = config.RUN_MODE_DEV
	if len(args) > 1 {
		args = args[1:]
		if len(args)%2 == 1 {
			log.Println("Incorrect  number of command line parameters", args)
			return cla
		}
		for i := 0; i < len(args); i += 2 {
			if args[i] == config.RUN_MODE {
				cla[config.RUN_MODE] = args[i+1]
			}
		}

	}

	return cla
}

func initDB() {
}

func RunTests() {
	baseContext, cancel := context.WithCancel(context.Background())
	signalChannel := registerShutdownHook(cancel)
	mainGroup, groupContext := errgroup.WithContext(baseContext)

	jq := queue.NeqJobQueue("test", mainGroup)
	jq.Start(groupContext)
	go func() {
		for i := 0; i < 10; i++ {
			log.Println(nonVoidFunc(jq, i))
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			voidFunc(jq, i)
		}
	}()

	if err := mainGroup.Wait(); err == nil {
		log.Println("FINISH CORE")
	}

	defer close(signalChannel)
}

func voidFunc(jq queue.IJobQueue, i int) {
	workIndex := i
	jq.Add(&queue.Job{Execute: func() {
		log.Println("Work ", workIndex, "Start")
		time.Sleep(time.Second * 4)
		log.Println("Work", workIndex, "Done")
	}})
}

func nonVoidFunc(jq queue.IJobQueue, i int) string {
	resChan := make(chan string)
	defer close(resChan)
	workIndex := i
	jq.Add(&queue.Job{Execute: func() {
		log.Println("WorkXX ", workIndex, "Start")
		time.Sleep(time.Second * 2)
		log.Println("WorkXX", workIndex, "Done")
		resChan <- "Result" + strconv.Itoa(workIndex)
	}})

	return <-resChan

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
