package handlers

import (
	"errors"
	"gopt/core/domain/project"
	"gopt/coreapi/viewhandlers"
	"log/slog"

	view_errors "gopt/public/error"
	view_project "gopt/public/project"

	"github.com/gin-gonic/gin"
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

func (h projectHandler) projectsPage(c *gin.Context) {
	slog.Info("PROJECTS PAGE")
	view_project.Projects(h.readRepo.GetProjects().Data()).Render(c.Request.Context(), c.Writer)
}

func (h projectHandler) newProject(c *gin.Context) {
	view_project.NewProject().Render(c.Request.Context(), c.Writer)
}

func (h projectHandler) addProject(c *gin.Context) {
	slog.Info("addProject")
	command := project.CreateProjectCommand{
		IssueKey: c.PostForm("project-key"),
		Name:     c.PostForm("project-name"),
	}
	err := validateProject(command)
	if err != nil {
		view_project.ProjectAddError(err.Error()).Render(c.Request.Context(), c.Writer)
		return
	}
	result := h.projectService.CreateProject(command)
	if !result.Sucess() {
		view_project.ProjectAddError(result.Error().Error()).Render(c.Request.Context(), c.Writer)
		return
	}
	slog.Info("Project Created")
	c.Writer.Header().Add("HX-Redirect", "/gopt/views/projects")
	// c.Redirect(http.StatusFound, "/gopt/projects")
}

func (h projectHandler) projectDetails(c *gin.Context) {
	projectId := c.Param("itemId")
	result := h.projectService.GetProject(projectId)
	if !result.Sucess() {
		view_errors.Error(result.Error().Error()).Render(c.Request.Context(), c.Writer)
		return
	}
	slog.Info("PROJECT DETAILS", "data", result.Data())
	view_project.ProjectDetails(result.Data()).Render(c.Request.Context(), c.Writer)
}

func validateProject(command project.CreateProjectCommand) error {
	var result string
	if len(command.Name) == 0 {
		result = "Name must not be empty.\n"
	}
	if len(command.IssueKey) == 0 {
		result += "Key must not be empty"
	}
	if len(result) != 0 {
		return errors.New(result)
	} else {
		return nil
	}
}
