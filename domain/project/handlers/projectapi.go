package handlers

import (
	"gosi/domain/project"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler projectHandler) getProjects(c *gin.Context) {
	log.Println("getProjetcs")
	c.JSON(http.StatusOK, gin.H{"data": handler.readRepo.GetProjects(), "error": ""})
}

func (handler projectHandler) getProject(c *gin.Context) {
	projectId := c.Param("itemId")

	result := handler.projectService.GetProject(projectId)
	if !result.Sucess() {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error(), "data": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result.Data(), "error": ""})
}

func (handler projectHandler) addProjectAPI(c *gin.Context) {
	var newProject project.CreateProjectCommand

	if err := c.BindJSON(&newProject); err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	result := handler.projectService.CreateProject(newProject)
	if !result.Sucess() {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error(), "data": data})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": result.Data(), "error": ""})
}
