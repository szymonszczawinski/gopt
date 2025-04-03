package handlers

import (
	"gopt/core/domain/common"
	"gopt/core/domain/issue"
	"gopt/core/domain/project"
	"gopt/coreapi"
	"log/slog"

	view_issue "gopt/public/issue"

	"github.com/gin-gonic/gin"
)

type IIssueRepo interface {
	GetIssue(key string) coreapi.Result[issue.Issue]
	GetIssues() coreapi.Result[[]issue.IssueListElement]
}
type IIssueCache interface {
	GetIssueTypes() []common.IssueType
	GetAvailabeProjects() []project.ProjectListElement
}
type issueHandler struct {
	repo  IIssueRepo
	cache IIssueCache
}

func NewIssueHandler(repo IIssueRepo, cache IIssueCache) issueHandler {
	instance := issueHandler{
		cache: cache,
		repo:  repo,
	}
	return instance
}

func (handler issueHandler) ConfigureRoutes(path string, routes Routes) {
	issuesRoute := routes.Views().Group(path)
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
	issueTypes := h.cache.GetIssueTypes()
	availableProjects := h.cache.GetAvailabeProjects()
	view_issue.NewIssue(issueTypes, availableProjects).Render(c.Request.Context(), c.Writer)
}

func (h issueHandler) addIssue(c *gin.Context) {
	slog.Info("ADD ISSUE")
	command, err := issue.NewCreateIssue(
		common.IssueType(c.PostForm("issue-type")),
		c.PostForm("issue-name"),
		c.PostForm("project-key"),
		c.PostForm("issue-summary"),
	)
	slog.Info("received command", "cmd", command, "err", err)
}

func (h issueHandler) listIssues(c *gin.Context) {
	slog.Info("ISSUE ALL")
	result := h.repo.GetIssues()
	if !result.Sucess() {
		slog.Error("get issue list", "err", result.Error().Error())
		return
	}
	isHxRequest := c.GetHeader("HX-Request")
	if isHxRequest == "true" {
		view_issue.Issues(true, result.Data()).Render(c.Request.Context(), c.Writer)
	} else {
		view_issue.Issues(false, result.Data()).Render(c.Request.Context(), c.Writer)
	}
}
