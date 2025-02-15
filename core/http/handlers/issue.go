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
		issueRoutes.GET("/", handler.issueAll)
		issueRoutes.GET("/:itemId", handler.issueDetails)
		issueRoutes.POST("/", handler.issueAdd)
	}
}

func (h issueHandler) issueDetails(c *gin.Context) {
	issueId := c.Param("itemId")
	slog.Info("ISSUE DETAILS", "issue id", issueId)
	view_issue.IssueDetails(issueId).Render(c.Request.Context(), c.Writer)
}

func (h issueHandler) issueAdd(c *gin.Context) {
	slog.Info("ISSUE ADD")
}

func (h issueHandler) issueAll(c *gin.Context) {
	slog.Info("ISSUE ALL")
}
