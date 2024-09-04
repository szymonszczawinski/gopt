package core

import (
	"context"
	"gopt/core/config"
	"gopt/core/domain/auth"
	"gopt/core/domain/home"
	"gopt/core/domain/project"
	"gopt/core/http"
	"gopt/core/messenger"
	"gopt/core/service"
	"gopt/core/storage/repository/postgres"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	repo_auth "gopt/core/storage/repository/postgres/auth"
	repo_project "gopt/core/storage/repository/postgres/project"

	project_handlers "gopt/core/domain/project/handlers"
	api_service "gopt/coreapi/service"

	"golang.org/x/sync/errgroup"
)

const (
	DefaultHttpServerPort = 8081
)

func Start(cla map[string]any, staticContent http.StaticContent) {
	slog.Info("START CORE")
	config.InitSystemConfiguration(cla)
	slog.Info("Start Parameters:", cla)
	baseContext, cancel := context.WithCancel(context.Background())
	signalChannel := registerShutdownHook(cancel)
	mainGroup, groupContext := errgroup.WithContext(baseContext)
	sm := service.NewServiceManager(mainGroup, groupContext)
	// some simple comment
	startServices(sm, mainGroup, groupContext, staticContent)
	// time.Sleep(time.Second * 5)
	if err := mainGroup.Wait(); err == nil {
		slog.Info("FINISH CORE")
	}

	defer close(signalChannel)
}

func startServices(sm api_service.IServiceManager, eg *errgroup.Group, ctx context.Context, staticContent http.StaticContent) {
	slog.Info("START CORE :: START SERVICES")

	slog.Info("Starting MESSENGER SERVICE")
	messengerService := messenger.NewMessengerService(eg, ctx)
	sm.StartComponent(api_service.ComponentTypeMessenger, messengerService)

	slog.Info("Starting DATABASE")
	// databaseConnection := bun.NewBunDatabase(eg, ctx)
	databaseConnection := postgres.NewPostgresSqlDatabase(eg, ctx)
	sm.StartComponent(api_service.ComponentTypeSqlDatabase, databaseConnection)

	slog.Info("Starting PROJECT REPOSITORY")
	// projectRepository := project.NewProjectRepositoryBun(eg, ctx, databaseConnection)
	projectRepository := repo_project.NewProjectRepositoryPostgres(eg, ctx, databaseConnection)
	sm.StartComponent(api_service.ComponentTypeProjectRepository, projectRepository)

	slog.Info("Starting PROJECT SERVICE")
	projetcsService := project.NewProjectService(eg, ctx, projectRepository)
	sm.StartComponent(api_service.ComponentTypeProjectService, projetcsService)

	slog.Info("Starting AUTH REPOSITORY")
	authRepository := repo_auth.NewAuthRepository(eg, ctx, databaseConnection)
	// sm.StartComponent(iservice.ComponentTypeAuthRepository, authRepository)

	slog.Info("Starting AUTH SERVICE")
	authService := auth.NewAuthenticationService(eg, ctx, authRepository)
	sm.StartComponent(api_service.ComponentTypeAuthService, authService)

	homeController := home.NewHomeHandler()
	projectsController := project_handlers.NewProjectHandler(projetcsService, projectRepository)
	authController := auth.NewAuthHandler(authService)

	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		httpPort = DefaultHttpServerPort
	}

	httpServer := http.NewHttpServer(ctx, eg, httpPort, staticContent)
	httpServer.AddHandler(homeController)
	httpServer.AddHandler(projectsController)
	httpServer.AddHandler(authController)
	httpServer.Start()

	// slog.Info("Starting HTTP SERVER SERVICE")
	// httpServerService := http.NewHttpServerService(eg, ctx, staticContent)
	// sm.StartComponent(iservice.ComponentTypeHttpServerService, httpServerService)
}

// func createPluginService(serviceLocation string, serviceName string) iservice.IComponent {
// 	plug, err := plugin.Open(serviceLocation)
// 	if err != nil {
// 		slog.Info("Could not load: ", serviceName, "Error: ", err)
// 		return nil
// 	}
// 	createMethod, err := plug.Lookup(iservice.NEW_FUNCTION)
// 	if err != nil {
// 		slog.Info("Could not get New from: ", serviceName)
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
// 		slog.Info("Instance is not IModService")
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
