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
type ProjectListElement struct {
	ProjectKey string `json:"projectKey"`
	Name       string `json:"name"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
	State      string `json:"state"`
	Owner      string `json:"owner"`
	Id         int    `json:"id"`
}

func NewProjectListElement(id int, projectKey, name, projectState, owner string, projectCreationTime, projectLastUpdateTime time.Time) ProjectListElement {
	return ProjectListElement{
		Id:         id,
		ProjectKey: projectKey,
		Name:       name,
		State:      projectState,
		Owner:      owner,
		Created:    projectCreationTime.Format(DDMMYYYYhhmmss),
		Updated:    projectLastUpdateTime.Format(DDMMYYYYhhmmss),
	}
}

// Read-only view of a project
type ProjectDetails struct {
	ProjectKey string `json:"projectKey"`
	Name       string `json:"name"`
	State      string `json:"state"`
	Owner      string `json:"owner"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
	Items      []ProjectDetailsItem
	Id         int `json:"id"`
}

func NewProjectDetails(project Project) ProjectDetails {
	projectDetails := ProjectDetails{
		Id:         project.GetId(),
		ProjectKey: project.projectKey,
		Name:       project.name,
		State:      project.state.String(),
		Owner:      project.GetOwner(),
		Created:    project.GetCreationTime().String(),
		Updated:    project.GetLastUpdateTime().String(),
	}

	return projectDetails
}

// Read-only view of project related items: Tasks, Bugs, etc
type ProjectDetailsItem struct {
	ItemType string
	Name     string
	ItemKey  string
	State    string
}

type ProjectComment struct {
	Created time.Time `json:"created"`
	Content string    `json:"content"`
	Id      int       `json:"id"`
}
