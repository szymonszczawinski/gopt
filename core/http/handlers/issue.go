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
	issuesRoute := routes.Views().Group("/issues")
	// pagesProjects.Use(auth.SessionAuth)

	issuesRoute.GET("/", handler.listIssues)
	issuesRoute.GET("/:itemId", handler.issueDetails)
	issuesRoute.GET("/new", handler.newIssue)
	issuesRoute.POST("/new", handler.addIssue)
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
	command, err := issue.NewCreateIssue(c.PostForm("project-key"), c.PostForm("issue-name"), c.PostForm("issue-type"))
	slog.Info("received command", "cmd", command, "err", err)
}

func (h issueHandler) listIssues(c *gin.Context) {
	slog.Info("ISSUE ALL")
}
