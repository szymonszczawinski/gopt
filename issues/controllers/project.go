package controllers

import (
	"gosi/issues/service"

	"github.com/gin-gonic/gin"
)

var projectService *service.ProjectService

func AddProjectsRoutes(apiRootRoute *gin.RouterGroup, rootRoute *gin.RouterGroup) {
	projectService = service.NewProjectService()
	apiRoute := apiRootRoute.Group("/projects")
	apiRoute.GET("/", getProjects)
	apiRoute.GET("/:issueId", getProject)
	apiRoute.POST("/add", addProjectAPI)
	apiRoute.POST("/:issueId/addComment", addProjectComment)

	projectsRoute := rootRoute.Group("/projects")
	projectsRoute.GET("/", projectsPage)
	projectsRoute.POST("/add", addProject)
}
