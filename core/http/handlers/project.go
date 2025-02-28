package handlers

import (
	"fmt"
	"gopt/core/domain/project"
	"gopt/coreapi"
	"log/slog"

	view_errors "gopt/public/error"
	view_project "gopt/public/project"

	"github.com/gin-gonic/gin"
)

type IProjectQueryRepository interface {
	coreapi.IComponent
	GetProjects() coreapi.Result[[]project.ProjectListElement]
}

type IProjectService interface {
	// GetProjects() coreapi.Result[[]ProjectListElement]
	GetProject(command project.GetProject) coreapi.Result[project.ProjectDetails]
	CreateProject(command project.CreateProject) coreapi.Result[project.ProjectDetails]
	CloseProject(projectId string) coreapi.Result[project.ProjectDetails]
}

type projectHandler struct {
	projectService IProjectService
	readRepo       IProjectQueryRepository
}

func NewProjectHandler(projectService IProjectService, readRepo IProjectQueryRepository) *projectHandler {
	instance := projectHandler{
		projectService: projectService,
		readRepo:       readRepo,
	}
	return &instance
}

func (handler *projectHandler) ConfigureRoutes(routes Routes) {
	projectsRoute := routes.Views().Group("/projects")
	// pagesProjects.Use(auth.SessionAuth)

	projectsRoute.GET("/", handler.listProjects)
	projectsRoute.GET("/:item_key", handler.projectDetails)
	projectsRoute.GET("/new", handler.newProject)
	projectsRoute.POST("/new", handler.addProject)
}

func (h projectHandler) listProjects(c *gin.Context) {
	slog.Info("PROJECTS PAGE")
	isHxRequest := c.GetHeader("HX-Request")
	if isHxRequest == "true" {
		view_project.Projects(true, h.readRepo.GetProjects().Data()).Render(c.Request.Context(), c.Writer)
	} else {
		view_project.Projects(false, h.readRepo.GetProjects().Data()).Render(c.Request.Context(), c.Writer)
	}
}

func (h projectHandler) newProject(c *gin.Context) {
	view_project.NewProject().Render(c.Request.Context(), c.Writer)
}

func (h projectHandler) addProject(c *gin.Context) {
	slog.Info("addProject")
	command, err := project.NewCreateProject(c.PostForm("project_key"), c.PostForm("project_name"))
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
	c.Writer.Header().Add("HX-Redirect", fmt.Sprintf("/gopt/views/projects/%v", result.Data().ProjectKey))
	// c.Redirect(http.StatusFound, "/gopt/projects")
}

func (h projectHandler) projectDetails(c *gin.Context) {
	command, err := project.NewGetProject(c.Param("item_key"))
	if err != nil {
		slog.Error("get project details", "err", err)
		// view_project.ProjectAddError(err.Error()).Render(c.Request.Context(), c.Writer)
		return
	}

	result := h.projectService.GetProject(command)
	if !result.Sucess() {
		view_errors.Error(result.Error().Error()).Render(c.Request.Context(), c.Writer)
		return
	}
	slog.Info("PROJECT DETAILS", "data", result.Data())
	isHxRequest := c.GetHeader("HX-Request")
	if isHxRequest == "true" {
		view_project.ProjectDetails(true, result.Data()).Render(c.Request.Context(), c.Writer)
	} else {
		view_project.ProjectDetails(false, result.Data()).Render(c.Request.Context(), c.Writer)
	}
}
