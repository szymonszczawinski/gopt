package handlers

import (
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
		issueRoutes.GET("/:itemId", handler.issueDetails)
	}
}

func (h issueHandler) issueDetails(c *gin.Context) {
	issueId := c.Param("itemId")
	slog.Info("ISSUE DETAILS", "issue id", issueId)
	view_issue.IssueDetails(issueId).Render(c.Request.Context(), c.Writer)
}
