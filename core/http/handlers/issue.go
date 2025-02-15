package handlers

import (
	"gopt/core/domain/issue"
	"gopt/coreapi/viewhandlers"
	"log/slog"

	view_issue "gopt/public/issue"

	"github.com/gin-gonic/gin"
)

type issueHandler struct{}

func NewIssueHandler() *issueHandler {
	instance := issueHandler{}
	return &instance
}

func (handler *issueHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	issueRoutes := routes.Views().Group("/issues")
	// pagesProjects.Use(auth.SessionAuth)
	{
		issueRoutes.GET("/", handler.allIssues)
		issueRoutes.GET("/new", handler.newIssue)
		issueRoutes.GET("/:itemId", handler.issueDetails)
		issueRoutes.POST("/new", handler.addIssue)
	}
}

func (h issueHandler) issueDetails(c *gin.Context) {
	issueId := c.Param("itemId")
	slog.Info("ISSUE DETAILS", "issue id", issueId)
	view_issue.IssueDetails(issueId).Render(c.Request.Context(), c.Writer)
}

func (h issueHandler) newIssue(c *gin.Context) {
	slog.Info("NEW ISSUE")
	view_issue.NewIssue().Render(c.Request.Context(), c.Writer)
}

func (h issueHandler) addIssue(c *gin.Context) {
	slog.Info("ADD ISSUE")
	command := issue.CreateIssueCommand{
		ProjectKey: c.PostForm("project-key"),
		Name:       c.PostForm("issue-name"),
		IssueType:  c.PostForm("issue-type"),
	}
	slog.Info("received command", "cmd", command)
}

func (h issueHandler) allIssues(c *gin.Context) {
	slog.Info("ISSUE ALL")
}
