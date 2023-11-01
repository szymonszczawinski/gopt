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

func (self projectHandler) projectsPage(c *gin.Context) {
	log.Println("PROJECTS PAGE")
	c.HTML(http.StatusOK, "projects", gin.H{
		"title": "Projects",
		"data":  self.projectService.GetProjects(), "error": ""})
}
func (self projectHandler) newProject(c *gin.Context) {
	c.HTML(http.StatusOK, "projects/newproject.html", gin.H{"title": "Add Project"})
}

func (self projectHandler) addProject(c *gin.Context) {
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
	_, err = self.projectService.CreateProject(newProject)
	if err != nil {
		displayeError(err, newProject, c)
		return
	}
	log.Println("Project Created")
	c.Writer.Header().Add("HX-Redirect", "/gosi/projects")
	// c.Redirect(http.StatusFound, "/gosi/projects")
}

func (self projectHandler) projectDetails(c *gin.Context) {
	projectId := c.Param("issueId")
	project, err := self.projectService.GetProject(projectId)
	if err != nil {
		log.Println(err.Error())
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	log.Println("PROJECT DETAILS", project)
	c.HTML(http.StatusOK, "project_details", gin.H{
		"title": "Project Details",
		"data":  project, "error": ""})

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
