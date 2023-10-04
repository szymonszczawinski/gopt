package handlers

import (
	"embed"
	// "gosi/auth"
	"gosi/coreapi/viewhandlers"
	"gosi/project"

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

func (self *projectHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	pagesProjects := routes.Views().Group("/projects")
	// pagesProjects.Use(auth.SessionAuth)
	{
		pagesProjects.GET("/", self.projectsPage)
		pagesProjects.GET("/new", self.newProject)
		pagesProjects.POST("/new", self.addProject)
		pagesProjects.GET("/:issueId", self.projectDetails)
	}
}
func (self *projectHandler) ConfigureApiRoutes(routes viewhandlers.Routes) {
	apiProjects := routes.Apis().Group("/project")

	apiProjects.GET("/", self.getProjects)
	apiProjects.GET("/:issueId", self.getProject)
	apiProjects.POST("/add", self.addProjectAPI)

}

func (self *projectHandler) LoadViews(r multitemplate.Renderer) multitemplate.Renderer {
	viewhandlers.AddCompositeView(r, "projects", "public/project/projects.html", viewhandlers.GetLayouts(), self.FileSystem)
	viewhandlers.AddCompositeView(r, "project_details", "public/project/details.html", viewhandlers.GetLayouts(), self.FileSystem)
	return r
}
