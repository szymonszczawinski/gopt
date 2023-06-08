package issues

import (
	"gosi/users"
	"time"
)

type IssueType int
type RelationType int

const (
	TProject IssueType = 1
	TBug     IssueType = 2
)

const (
	RCauses     RelationType = 1
	RIsCausedBy RelationType = 2
)

type Issue struct {
	Id           int64          `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	CurrentState LifecycleState `json:"state"`
	Lifecycle    Lifecycle      `json:"lifecycle"`
	created      time.Time
	updated      time.Time
	comments     []Comment
	relations    []Relation
}

type Project struct {
	Issue
	IssueType IssueType `json:"issueType"`
}

type Bug struct {
	Issue
	IssueType IssueType
}

type Relation struct {
	id           int64
	relationType RelationType
	fromIssue    Issue
	toIssue      Issue
}

type Comment struct {
	Id      int64
	ItemId  int64
	Author  users.User
	created time.Time
}

func NewProject(id int64, name string, lifecycle Lifecycle) *Project {
	project := Project{
		Issue: Issue{
			Id:           id,
			Name:         name,
			Description:  "",
			CurrentState: lifecycle.startState,
			Lifecycle:    lifecycle,
		},
		IssueType: TProject,
	}
	return &project
}
