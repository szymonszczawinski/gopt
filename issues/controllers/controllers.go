package controllers

import (
	"embed"
	"gosi/coreapi/viewcon"
	projectservice "gosi/issues/service"

	"github.com/gin-gonic/gin"
)

type projectController struct {
	viewcon.Controller
	projectService projectservice.IProjectService
}

func NewProjectController(projectService projectservice.IProjectService, fs embed.FS) *projectController {
	instance := projectController{
		Controller: viewcon.Controller{
			FileSystem: fs,
		},
		projectService: projectService,
	}
	return &instance
}

func (self *projectController) ConfigureRoutes(root, pages, api *gin.RouterGroup, fs embed.FS) {
	apiProjects := api.Group("/projects")

	apiProjects.GET("/", self.getProjects)
	apiProjects.GET("/:issueId", self.getProject)
	apiProjects.POST("/add", self.addProjectAPI)
	apiProjects.POST("/:issueId/addComment", self.addProjectComment)

	pagesProjects := pages.Group("/projects")

	pagesProjects.GET("/", self.projectsPage)
	pagesProjects.GET("/new", self.newProject)
	pagesProjects.POST("/new", self.addProject)
	pagesProjects.GET("/:issueId", self.projectDetails)
}

func AddProjectsRoutes(apiRootRoute *gin.RouterGroup, rootRoute *gin.RouterGroup) {

}
