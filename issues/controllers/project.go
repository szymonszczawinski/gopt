package controllers

import (
	"errors"
	"gosi/core/service"
	iservice "gosi/coreapi/service"
	"gosi/coreapi/storage"
	projectservice "gosi/issues/service"
	"log"

	"github.com/gin-gonic/gin"
)

var projectService *projectservice.ProjectService

func AddProjectsRoutes(apiRootRoute *gin.RouterGroup, rootRoute *gin.RouterGroup) {

	storageService, _ := getStorageService()
	repository, _ := getRepository()
	projectService = projectservice.NewProjectService(storageService, repository)
	apiProjectsRoute := apiRootRoute.Group("/projects")
	apiProjectsRoute.GET("/", getProjects)
	apiProjectsRoute.GET("/:issueId", getProject)
	apiProjectsRoute.POST("/add", addProjectAPI)
	apiProjectsRoute.POST("/:issueId/addComment", addProjectComment)

	projectsRoute := rootRoute.Group("/projects")
	projectsRoute.GET("/", projectsPage)
	projectsRoute.GET("/new", newProject)
	projectsRoute.POST("/new", addProject)
}

func getStorageService() (storage.IStorageService, error) {

	serviceManager, err := service.GetServiceManager()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	service, err := serviceManager.GetService(iservice.ServiceTypeIStorageService)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	storageService, ok := service.(storage.IStorageService)
	if !ok {
		log.Fatal(err.Error())
		return nil, errors.New("StorageService has incorrect type")
	}
	return storageService, nil
}
func getRepository() (storage.IRepository, error) {

	serviceManager, err := service.GetServiceManager()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	service, err := serviceManager.GetService(iservice.ServiceTypeIRepository)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	repository, ok := service.(storage.IRepository)
	if !ok {
		log.Fatal(err.Error())
		return nil, errors.New("Repository has incorrect type")
	}
	return repository, nil
}
