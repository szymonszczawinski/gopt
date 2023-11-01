package project

import (
	"time"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

type CreateProjectCommand struct {
	IssueKey string `json:"issueKey"`
	Name     string `json:"name"`
}

type AddCommentCommand struct {
	ParentIssueKey string `json:"parrentIssueKey"`
	Content        string `json:"content"`
}
type ProjectListItem struct {
	Id         int    `json:"id"`
	ProjectKey string `json:"projectKey"`
	Name       string `json:"name"`
	State      string `json:"state"`
	Owner      string `json:"owner"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

type ProjectDetails struct {
	Id         int    `json:"id"`
	ProjectKey string `json:"projectKey"`
	Name       string `json:"name"`
	State      string `json:"state"`
	Owner      string `json:"owner"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

type ProjectComment struct {
	Id      int       `json:"id"`
	Created time.Time `json:"created"`
	Content string    `json:"content"`
}

func NewProjectListItem(project Project) ProjectListItem {
	return ProjectListItem{
		Id:         project.GetId(),
		ProjectKey: project.ProjectKey,
		Name:       project.Name,
		State:      project.State.String(),
		Owner:      "",
		Created:    project.GetCreationTime().Format(DDMMYYYYhhmmss),
		Updated:    project.GetLastUpdateTime().Format(DDMMYYYYhhmmss),
	}
}
func NewProjectDetails(project Project) ProjectDetails {
	projectDetails := ProjectDetails{
		Id:         project.GetId(),
		ProjectKey: project.ProjectKey,
		Name:       project.Name,
		State:      project.State.String(),
		Owner:      "",
		Created:    project.GetCreationTime().String(),
		Updated:    project.GetLastUpdateTime().String(),
	}

	return projectDetails

}
