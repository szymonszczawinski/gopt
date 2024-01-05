package handlers

import (
	"errors"
	"gosi/domain/project"
	"log"

	view_errors "gosi/public/error"
	view_project "gosi/public/project"

	"github.com/gin-gonic/gin"
)

func (h projectHandler) projectsPage(c *gin.Context) {
	log.Println("PROJECTS PAGE")
	view_project.Projects(h.readRepo.GetProjects().Data()).Render(c.Request.Context(), c.Writer)
}

func (h projectHandler) newProject(c *gin.Context) {
	view_project.NewProject().Render(c.Request.Context(), c.Writer)
}

func (h projectHandler) addProject(c *gin.Context) {
	log.Println("addProject")
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
	log.Println("Project Created")
	c.Writer.Header().Add("HX-Redirect", "/gosi/projects")
	// c.Redirect(http.StatusFound, "/gosi/projects")
}

func (h projectHandler) projectDetails(c *gin.Context) {
	projectId := c.Param("itemId")
	result := h.projectService.GetProject(projectId)
	if !result.Sucess() {
		log.Println("EEEEEEE", result.Error())
		view_errors.Error(result.Error().Error()).Render(c.Request.Context(), c.Writer)
		return
	}
	log.Println("PROJECT DETAILS", result.Data())
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
