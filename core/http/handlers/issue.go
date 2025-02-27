package handlers

import (
	"gopt/core/domain/issue"
	"gopt/coreapi"
	"gopt/coreapi/viewhandlers"
	"log/slog"

	view_issue "gopt/public/issue"

	"github.com/gin-gonic/gin"
)

type IIssueRepo interface {
	GetIssue(key string) coreapi.Result[issue.Issue]
}
type issueHandler struct {
	repo IIssueRepo
}

func NewIssueHandler(repo IIssueRepo) *issueHandler {
	instance := issueHandler{
		repo: repo,
	}
	return &instance
}

func (handler *issueHandler) ConfigureRoutes(routes viewhandlers.Routes) {
	issuesRoute := routes.Views().Group("/issues")
	// pagesProjects.Use(auth.SessionAuth)

	issuesRoute.GET("/", handler.listIssues)
	issuesRoute.GET("/:itemKey", handler.issueDetails)
	issuesRoute.GET("/new", handler.newIssue)
	issuesRoute.POST("/new", handler.addIssue)
}

func (h issueHandler) issueDetails(c *gin.Context) {
	command, err := issue.NewGetIssue(c.Param("itemKey"))
	if err != nil {
		slog.Error("get issue details", "cmd", command, "err", err)
		return
	}
	slog.Info("ISSUE DETAILS", "cmd", command)
	result := h.repo.GetIssue(command.IssueKey)
	if !result.Sucess() {
		slog.Error("get issue details", "cmd", command, "err", result.Error().Error())
		return
	}
	isHxRequest := c.GetHeader("HX-Request")
	if isHxRequest == "true" {
		view_issue.IssueDetails(true, result.Data().GetItemKey(), result.Data().ParentKey()).Render(c.Request.Context(), c.Writer)
	} else {
		view_issue.IssueDetails(false, result.Data().GetItemKey(), result.Data().ParentKey()).Render(c.Request.Context(), c.Writer)
	}
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
