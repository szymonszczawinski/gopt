package handlers

import (
	"embed"
	"gosi/coreapi/viewhandlers"
	"gosi/domain/project"

	"github.com/gin-contrib/multitemplate"
)

var (
	ProjectsView       = viewhandlers.View{Name: "projects", Template: "public/project/projects.html"}
	ProjectDetailsView = viewhandlers.View{Name: "project_details", Template: "public/project/details.html"}
	ProjectNewView     = viewhandlers.View{Name: "project_new", Template: "public/project/newproject.html"}
)

type projectHandler struct {
	viewhandlers.BaseHandler
	projectService project.IProjectService
	readRepo       project.IProjectQueryRepository
}

func NewProjectHandler(projectService project.IProjectService, readRepo project.IProjectQueryRepository, fs embed.FS) *projectHandler {
	instance := projectHandler{
		BaseHandler: viewhandlers.BaseHandler{
			FileSystem: fs,
		},
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
	viewhandlers.AddCompositeView(r, ProjectsView.Name, ProjectsView.Template, viewhandlers.GetLayouts(), handler.FileSystem)
	viewhandlers.AddCompositeView(r, ProjectDetailsView.Name, ProjectDetailsView.Template, viewhandlers.GetLayouts(), handler.FileSystem)
	viewhandlers.AddCompositeView(r, ProjectNewView.Name, ProjectNewView.Template, viewhandlers.GetLayouts(), handler.FileSystem)
	return r
}
