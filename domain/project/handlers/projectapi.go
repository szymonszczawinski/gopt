package handlers

import (
	"gosi/domain/project"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler projectHandler) getProjects(c *gin.Context) {
	log.Println("getProjetcs")
	c.JSON(http.StatusOK, gin.H{"data": handler.projectService.GetProjects(), "error": ""})
}

func (handler projectHandler) getProject(c *gin.Context) {
	projectId := c.Param("issueId")

	project, err := handler.projectService.GetProject(projectId)
	if err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": project, "error": ""})

}

func (handler projectHandler) addProjectAPI(c *gin.Context) {
	var newProject project.CreateProjectCommand

	if err := c.BindJSON(&newProject); err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	createdProject, err := handler.projectService.CreateProject(newProject)
	if err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": createdProject, "error": ""})

}
