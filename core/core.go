package core

import (
	"context"
	"gosi/core/http"
	"gosi/core/messenger"
	"gosi/core/service"
	"gosi/core/storage/sql"
	"gosi/domain/home"
	"gosi/domain/project"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	iservice "gosi/coreapi/service"
	project_handlers "gosi/domain/project/handlers"

	"golang.org/x/sync/errgroup"
)

const (
	DefaultHttpServerPort = 8081
)

var systemStartParameters map[string]any

func Start(cla map[string]any, staticContent http.StaticContent) {
	log.Println("START CORE")
	systemStartParameters = cla
	log.Println("Start Parameters:", systemStartParameters)
	baseContext, cancel := context.WithCancel(context.Background())
	signalChannel := registerShutdownHook(cancel)
	mainGroup, groupContext := errgroup.WithContext(baseContext)
	service.NewServiceManager(mainGroup, groupContext)
	// some simple comment
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
	// databaseConnection := bun.NewBunDatabase(eg, ctx)
	databaseConnection := sql.NewPostgresSqlDatabase(eg, ctx)
	sm.StartComponent(iservice.ComponentTypeSqlDatabase, databaseConnection)

	log.Println("Starting PROJECT REPOSITORY")
	// projectRepository := project.NewProjectRepositoryBun(eg, ctx, databaseConnection)
	projectRepository := project.NewProjectRepositorySql(eg, ctx, databaseConnection)
	sm.StartComponent(iservice.ComponentTypeProjectRepository, projectRepository)

	log.Println("Starting PROJECT SERVICE")
	projetcsService := project.NewProjectService(eg, ctx, projectRepository)
	sm.StartComponent(iservice.ComponentTypeProjectService, projetcsService)

	// log.Println("Starting AUTH REPOSITORY")
	// authRepository := auth.NewAuthRepository(eg, ctx, databaseConnection)
	// sm.StartComponent(iservice.ComponentTypeAuthRepository, authRepository)

	// log.Println("Starting AUTH SERVICE")
	// authService := auth.NewAuthenticationService(eg, ctx, authRepository)
	// sm.StartComponent(iservice.ComponentTypeAuthService, authService)

	homeController := home.NewHomeHandler(staticContent.PublicDir)
	projectsController := project_handlers.NewProjectHandler(projetcsService, projectRepository, staticContent.PublicDir)
	// authController := auth.NewAuthHandler(authService, staticContent.PublicDir)

	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		httpPort = DefaultHttpServerPort
	}

	httpServer := http.NewHttpServer(ctx, eg, httpPort, staticContent)
	httpServer.AddHandler(homeController)
	httpServer.AddHandler(projectsController)
	// httpServer.AddHandler(authController)
	httpServer.Start()

	// log.Println("Starting HTTP SERVER SERVICE")
	// httpServerService := http.NewHttpServerService(eg, ctx, staticContent)
	// sm.StartComponent(iservice.ComponentTypeHttpServerService, httpServerService)
}

// func createPluginService(serviceLocation string, serviceName string) iservice.IComponent {
// 	plug, err := plugin.Open(serviceLocation)
// 	if err != nil {
// 		log.Println("Could not load: ", serviceName, "Error: ", err)
// 		return nil
// 	}
// 	createMethod, err := plug.Lookup(iservice.NEW_FUNCTION)
// 	if err != nil {
// 		log.Println("Could not get New from: ", serviceName)
// 		return nil
// 	}
// 	createFunction, isCreateFunction := createMethod.(func() any)
// 	if !isCreateFunction {
// 		log.Printf("Not ceate function %T", createMethod)
// 		return nil
// 	}
// 	instance := createFunction()
// 	serviceInstance, isInstance := instance.(iservice.IComponent)
// 	if !isInstance {
// 		log.Println("Instance is not IModService")
// 		return nil
// 	}
// 	return serviceInstance
// }

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
