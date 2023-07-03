package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func projectsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "projects/index.html", gin.H{
		"title": "Projects",
		"data":  projectService.GetProjects(), "error": ""})
}
