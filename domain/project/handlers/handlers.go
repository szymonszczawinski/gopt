package handlers

import (
	"embed"
	"gosi/coreapi/viewhandlers"
	"gosi/domain/project"

	"github.com/gin-contrib/multitemplate"
)

type projectHandler struct {
	viewhandlers.BaseHandler
	projectService project.IProjectService
}

func NewProjectHandler(projectService project.IProjectService, fs embed.FS) *projectHandler {
	instance := projectHandler{
		BaseHandler: viewhandlers.BaseHandler{
			FileSystem: fs,
		},
		projectService: projectService,
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
		pagesProjects.GET("/:issueId", handler.projectDetails)
	}
}
func (handler *projectHandler) ConfigureApiRoutes(routes viewhandlers.Routes) {
	apiProjects := routes.Apis().Group("/project")

	apiProjects.GET("/", handler.getProjects)
	apiProjects.GET("/:issueId", handler.getProject)
	apiProjects.POST("/add", handler.addProjectAPI)

}

func (handler *projectHandler) LoadViews(r multitemplate.Renderer) multitemplate.Renderer {
	viewhandlers.AddCompositeView(r, "projects", "public/project/projects.html", viewhandlers.GetLayouts(), handler.FileSystem)
	viewhandlers.AddCompositeView(r, "project_details", "public/project/details.html", viewhandlers.GetLayouts(), handler.FileSystem)
	return r
}
