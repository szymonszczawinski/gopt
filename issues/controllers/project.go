package controllers

import (
	"gosi/issues/dto"
	"gosi/issues/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var projectService service.ProjectService

func AddProjectsRoutes(apiRootRoute *gin.RouterGroup, rootRoute *gin.RouterGroup) {
	projectService = service.NewProjectService()

	apiRoute := apiRootRoute.Group("/projects")
	apiRoute.GET("/", getProjects)
	apiRoute.GET("/:issueId", getProject)
	apiRoute.POST("/add", addProject)
	apiRoute.POST("/:issueId/addComment", addProjectComment)

	projectsRoute := rootRoute.Group("/projects")
	projectsRoute.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects/index.html", gin.H{
			"title": "PROJECTS",
		})
	})
}

func getProjects(c *gin.Context) {
	log.Println("getProjetcs")
	c.JSON(http.StatusOK, gin.H{"data": projectService.GetProjects(), "error": ""})
}

func getProject(c *gin.Context) {
	projectId := c.Param("issueId")

	project, err := projectService.GetProject(projectId)
	if err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": project, "error": ""})

}

func addProject(c *gin.Context) {
	var newProject dto.CreateProjectCommand

	if err := c.BindJSON(&newProject); err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	createdProject, err := projectService.CreateProject(newProject)
	if err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"data": createdProject, "error": ""})

}

func addProjectComment(c *gin.Context) {
	var projectComment dto.AddCommentCommand
	if err := c.BindJSON(&projectComment); err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
	}
	createdComment, err := projectService.AddComment(projectComment)
	if err != nil {
		data := make(map[string]string)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": data})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"elementId": createdComment.GetId(), "error": ""})
}
