package controllers

import (
	"gosi/core/service"
	iservice "gosi/coreapi/service"
	projectservice "gosi/issues/service"
	"log"

	"github.com/gin-gonic/gin"
)

var projectService *projectservice.ProjectService

func AddProjectsRoutes(apiRootRoute *gin.RouterGroup, rootRoute *gin.RouterGroup) {

	projectService = mustCreateProjectService()
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

func mustCreateProjectService() *projectservice.ProjectService {

	serviceManager, err := service.GetServiceManager()
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	service := serviceManager.MustGetComponent(iservice.ComponentTypeIssueService)
	projectService, ok := service.(*projectservice.ProjectService)
	if !ok {
		log.Fatal(err.Error())
		return nil
	}
	return projectService
}
