package core

import (
	"gosi/auth"
	"gosi/core/http"
	"gosi/core/messenger"
	"gosi/core/storage/bun"
	"gosi/issues/storage"

	"gosi/core/service"
	iservice "gosi/coreapi/service"
	projects_controller "gosi/issues/controllers"
	project_service "gosi/issues/service"

	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"plugin"
	"syscall"

	"golang.org/x/sync/errgroup"
)

const (
	HttpServerPort = 8081
)

var systemStartParameters map[string]any

func Start(cla map[string]any, staticContent http.StaticContent) {
	log.Println("START CORE")
	systemStartParameters = cla
	baseContext, cancel := context.WithCancel(context.Background())
	signalChannel := registerShutdownHook(cancel)
	mainGroup, groupContext := errgroup.WithContext(baseContext)
	service.NewServiceManager(mainGroup, groupContext)
	//some simple comment
	startServices(mainGroup, groupContext, staticContent)
	// time.Sleep(time.Second * 5)
	if err := mainGroup.Wait(); err == nil {
		log.Println("FINISH CORE")
	}

	defer close(signalChannel)
}

func startServices(eg *errgroup.Group, ctx context.Context, staticContent http.StaticContent) {
	log.Println("START CORE :: START SERVICES")

	startCoreServices(eg, ctx, staticContent)
}

func startCoreServices(eg *errgroup.Group, ctx context.Context, staticContent http.StaticContent) {
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
	projetcsService := project_service.NewProjectService(eg, ctx, issueRepository)
	sm.StartComponent(iservice.ComponentTypeIssueService, projetcsService)

	log.Println("Starting AUTH REPOSITORY")
	authRepository := auth.NewAuthRepository(eg, ctx, databaseConnection)
	sm.StartComponent(iservice.ComponentTypeAuthRepository, authRepository)

	log.Println("Starting AUTH SERVICE")
	authService := auth.NewAuthenticationService(eg, ctx, authRepository)
	sm.StartComponent(iservice.ComponentTypeAuthService, authService)

	homeController := http.NewHomeController(staticContent.PublicDir)
	projectsController := projects_controller.NewProjectController(projetcsService, staticContent.PublicDir)
	authController := auth.NewAuthController(authService, staticContent.PublicDir)

	httpServer := http.NewHttpServer(ctx, eg, HttpServerPort, staticContent)
	httpServer.AddController(homeController)
	httpServer.AddController(projectsController)
	httpServer.AddController(authController)
	httpServer.Start()

	// log.Println("Starting HTTP SERVER SERVICE")
	// httpServerService := http.NewHttpServerService(eg, ctx, staticContent)
	// sm.StartComponent(iservice.ComponentTypeHttpServerService, httpServerService)

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
