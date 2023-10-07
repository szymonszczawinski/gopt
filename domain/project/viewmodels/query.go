package dto

import (
	"gosi/domain/project/domain"
	"time"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

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

func NewProjectListItem(project domain.Project) ProjectListItem {
	return ProjectListItem{
		Id:         project.GetId(),
		ProjectKey: project.ProjectKey,
		Name:       project.Name,
		State:      project.GetState().GetValue(),
		Owner:      "",
		Created:    project.GetCreationTime().Format(DDMMYYYYhhmmss),
		Updated:    project.GetLastUpdateTime().Format(DDMMYYYYhhmmss),
	}
}
func NewProjectDetails(project domain.Project) ProjectDetails {
	projectDetails := ProjectDetails{
		Id:         project.GetId(),
		ProjectKey: project.ProjectKey,
		Name:       project.Name,
		State:      project.GetState().GetValue(),
		Owner:      "",
		Created:    project.GetCreationTime().String(),
		Updated:    project.GetLastUpdateTime().String(),
	}

	return projectDetails

}
