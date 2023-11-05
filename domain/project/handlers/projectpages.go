package handlers

import (
	"errors"
	"fmt"
	"gosi/domain/project"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func (handler projectHandler) projectsPage(c *gin.Context) {
	log.Println("PROJECTS PAGE")
	c.HTML(http.StatusOK, "projects", gin.H{
		"title": "Projects",
		"data":  handler.projectService.GetProjects(), "error": ""})
}
func (handler projectHandler) newProject(c *gin.Context) {
	c.HTML(http.StatusOK, "projects/newproject.html", gin.H{"title": "Add Project"})
}

func (handler projectHandler) addProject(c *gin.Context) {
	log.Println("addProject")
	newProject := project.CreateProjectCommand{
		IssueKey: c.PostForm("project-key"),
		Name:     c.PostForm("project-name"),
	}
	err := validateProject(newProject)
	if err != nil {
		displayeError2(err, c)
		return
	}
	result := handler.projectService.CreateProject(newProject)
	if !result.Sucess() {
		displayeError(result.Error(), newProject, c)
		return
	}
	log.Println("Project Created")
	c.Writer.Header().Add("HX-Redirect", "/gosi/projects")
	// c.Redirect(http.StatusFound, "/gosi/projects")
}

func (handler projectHandler) projectDetails(c *gin.Context) {
	projectId := c.Param("issueId")
	result := handler.projectService.GetProject(projectId)
	if !result.Sucess() {
		log.Println(result.Error())
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error(), "data": data})
		return
	}
	log.Println("PROJECT DETAILS", result.Data())
	c.HTML(http.StatusOK, "project_details", gin.H{
		"title": "Project Details",
		"data":  result.Data(), "error": ""})

}

func validateProject(p project.CreateProjectCommand) error {
	var result string
	if len(p.Name) == 0 {
		result = "Name must not be empty.\n"
	}
	if len(p.IssueKey) == 0 {
		result += "Key must not be empty"
	}
	if len(result) != 0 {
		return errors.New(result)
	} else {
		return nil
	}
}

func displayeError(err error, p project.CreateProjectCommand, c *gin.Context) {
	c.HTML(http.StatusBadRequest, "projects/newproject.html",
		gin.H{"title": "Add Project", "error": err.Error(), "projectName": p.Name, "projectKey": p.IssueKey})

}
func displayeError2(err error, c *gin.Context) {
	log.Println("Could not create a Project: ", err.Error())
	tmpl := template.Must(template.ParseFiles("public/projects/newproject.html"))
	tmpl.ExecuteTemplate(c.Writer, "create-project-error", gin.H{"error": fmt.Sprintf("Could not create a project: %v", err.Error())})
}
