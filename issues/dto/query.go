package dto

import (
	"gosi/issues/domain"
	"strconv"
	"time"
)

type ProjectListItem struct {
	Id       int    `json:"id"`
	IssueId  string `json:"issueId"`
	Name     string `json:"name"`
	Type     string `json:"issueType"`
	State    string `json:"state"`
	Assignee string `json:"assignee"`
	Reporter string `json:"reporter"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type ProjectDetails struct {
	Id       int    `json:"id"`
	IssueId  string `json:"issueId"`
	Name     string `json:"name"`
	Type     string `json:"issueType"`
	State    string `json:"state"`
	Assignee string `json:"assignee"`
	Reporter string `json:"reporter"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	Comments []ProjectComment
}

type ProjectComment struct {
	Id      int       `json:"id"`
	Created time.Time `json:"created"`
	Content string    `json:"content"`
}

func NewProjectListItem(project domain.Project) ProjectListItem {
	return ProjectListItem{
		Id:       project.GetId(),
		IssueId:  project.GetItemKey() + "-" + strconv.Itoa(project.GetItemNumber()),
		Name:     project.GetName(),
		Type:     project.GetIssueType().String(),
		State:    "",
		Assignee: "",
		Reporter: "",
		Created:  project.GetCreationTime().String(),
		Updated:  project.GetLastUpdateTime().String(),
	}
}
func NewProjectDetails(project domain.Project) ProjectDetails {
	projectDetails := ProjectDetails{
		Id:       project.GetId(),
		IssueId:  project.GetItemKey(),
		Name:     project.GetName(),
		Type:     project.GetIssueType().String(),
		State:    project.GetLifecycleState().GetValue(),
		Assignee: "",
		Reporter: "",
		Created:  project.GetCreationTime().String(),
		Updated:  project.GetLastUpdateTime().String(),
		Comments: []ProjectComment{},
	}

	for _, comment := range project.GetComments() {
		projectDetails.Comments = append(projectDetails.Comments, ProjectComment{
			Id:      comment.GetId(),
			Created: comment.GetCreationTime(),
			Content: comment.GetContent(),
		})
	}
	return projectDetails

}
