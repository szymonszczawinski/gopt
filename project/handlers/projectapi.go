package handlers

import (
	"gosi/project/viewmodels"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (self projectHandler) getProjects(c *gin.Context) {
	log.Println("getProjetcs")
	c.JSON(http.StatusOK, gin.H{"data": self.projectService.GetProjects(), "error": ""})
}

func (self projectHandler) getProject(c *gin.Context) {
	projectId := c.Param("issueId")

	project, err := self.projectService.GetProject(projectId)
	if err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": project, "error": ""})

}

func (self projectHandler) addProjectAPI(c *gin.Context) {
	var newProject dto.CreateProjectCommand

	if err := c.BindJSON(&newProject); err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	createdProject, err := self.projectService.CreateProject(newProject)
	if err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": createdProject, "error": ""})

}
