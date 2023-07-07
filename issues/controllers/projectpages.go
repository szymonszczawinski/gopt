package controllers

import (
	"errors"
	"fmt"
	"gosi/issues/dto"
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func projectsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "projects/index.html", gin.H{
		"title": "Projects",
		"data":  projectService.GetProjects(), "error": ""})
}
func newProject(c *gin.Context) {
	c.HTML(http.StatusOK, "projects/newproject.html", gin.H{"title": "Add Project"})
}

func addProject(c *gin.Context) {
	log.Println("addProject")
	newProject := dto.CreateProjectCommand{
		IssueKey: c.PostForm("project-key"),
		Name:     c.PostForm("project-name"),
	}
	err := validateProject(newProject)
	if err != nil {
		displayeError2(err, c)
		return
	}
	_, err = projectService.CreateProject(newProject)
	if err != nil {
		displayeError(err, newProject, c)
		return
	}
	log.Println("Project Created")
	c.Writer.Header().Add("HX-Redirect", "/gosi/projects")
	// c.Redirect(http.StatusFound, "/gosi/projects")

}

func validateProject(p dto.CreateProjectCommand) error {
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

func displayeError(err error, p dto.CreateProjectCommand, c *gin.Context) {
	c.HTML(http.StatusBadRequest, "projects/newproject.html",
		gin.H{"title": "Add Project", "error": err.Error(), "projectName": p.Name, "projectKey": p.IssueKey})

}
func displayeError2(err error, c *gin.Context) {
	log.Println("Could not create a Project: ", err.Error())
	tmpl := template.Must(template.ParseFiles("public/projects/newproject.html"))
	tmpl.ExecuteTemplate(c.Writer, "create-project-error", gin.H{"error": fmt.Sprintf("Could not create a project: %v", err.Error())})

}
