package handlers

import (
	"gopt/core/domain/project"
	"gopt/coreapi/viewhandlers"

	"github.com/gin-contrib/multitemplate"
)

type projectHandler struct {
	projectService project.IProjectService
	readRepo       project.IProjectQueryRepository
}

func NewProjectHandler(projectService project.IProjectService, readRepo project.IProjectQueryRepository) *projectHandler {
	instance := projectHandler{
		projectService: projectService,
		readRepo:       readRepo,
	}
	return &instance
}

func (handler *projectHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	pagesProjects := routes.Views().Group("/projects")
	// pagesProjects.Use(auth.SessionAuth)
	{
		pagesProjects.GET("/", handler.projectsPage)
		pagesProjects.GET("/new", handler.newProject)
		pagesProjects.POST("/new", handler.addProject)
		pagesProjects.GET("/:itemId", handler.projectDetails)
	}
}

func (handler *projectHandler) ConfigureApiRoutes(routes viewhandlers.Routes) {
	apiProjects := routes.Apis().Group("/project")

	apiProjects.GET("/", handler.getProjects)
	apiProjects.GET("/:itemId", handler.getProject)
	apiProjects.POST("/add", handler.addProjectAPI)
}

func (handler *projectHandler) LoadViews(r multitemplate.Renderer) multitemplate.Renderer {
	return r
}
