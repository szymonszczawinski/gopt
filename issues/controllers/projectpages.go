package controllers

import (
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
func addProject(c *gin.Context) {
	newProject := dto.CreateProjectCommand{
		IssueKey: c.PostForm("project-name"),
		Name:     c.PostForm("project-key"),
	}
	createdProject, err := projectService.CreateProject(newProject)
	if err != nil {
		data := make(map[string]string)
		log.Println("Could not create a Project: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	tmpl := template.Must(template.ParseFiles("public/projects/index.html"))
	tmpl.ExecuteTemplate(c.Writer, "project-list-element", createdProject)
}
