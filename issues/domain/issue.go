package domain

import (
	"fmt"
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
	Entity
	TimeTracked
	itemKey      string
	itemNumber   int
	name         string
	description  string
	issueType    IssueType
	currentState LifecycleState
	lifecycle    Lifecycle
	comments     []Comment
	relations    []Relation
}

func (self Issue) GetItemKey() string {
	return self.itemKey
}
func (self Issue) GetItemNumber() int {
	return self.itemNumber
}

func (self Issue) GetName() string {
	return self.name
}

func (self Issue) GetDescription() string {
	return self.description
}

func (self Issue) GetIssueType() IssueType {
	return self.issueType
}

func (self IssueType) String() string {
	return [...]string{"", "Project", "Bug"}[self]
}

func (self *Issue) GetAndIncrementItemNumber() int {
	self.itemNumber += 1
	return self.itemNumber
}

func (self Issue) GetLifecycleState() LifecycleState {
	return self.currentState
}
func (self *Issue) AddComment(comment Comment) {
	self.comments = append(self.comments, comment)
}

type Project struct {
	Issue
}

func (self Project) String() string {
	return fmt.Sprintf("Project[id:%v; key:%v; name:%v; state:%v\n comments:%v]",
		self.GetId(), self.GetItemKey(), self.GetName(), self.currentState, self.comments)
}

func (self Project) GetComments() []Comment {
	return self.comments
}

type Requirement struct {
	Issue
	project   Project
	issueType IssueType
}

type Bug struct {
	Issue
	requirement Requirement
}

type Relation struct {
	Entity
	relationType RelationType
	fromIssue    Issue
	toIssue      Issue
}

func NewProject(projectKey string, name string, lifecycle Lifecycle) Project {
	project := Project{
		Issue: Issue{
			Entity: Entity{},
			TimeTracked: TimeTracked{
				created: time.Now(),
				updated: time.Now(),
			},
			itemKey:      projectKey + "-1",
			itemNumber:   1,
			name:         name,
			description:  "",
			currentState: lifecycle.startState,
			lifecycle:    lifecycle,
			issueType:    TProject,
			comments:     []Comment{},
		},
	}
	return project
}
func NewProjectFromRepo(id int, created time.Time, updated time.Time, itemKey string, itemNumber int, name string,
	description string, state LifecycleState, lifecycle Lifecycle) Project {
	project := Project{
		Issue: Issue{
			Entity: Entity{
				id: id,
			},
			TimeTracked: TimeTracked{
				created: created,
				updated: updated,
			},
			itemKey:      itemKey,
			itemNumber:   itemNumber,
			name:         name,
			description:  description,
			issueType:    TProject,
			currentState: state,
			lifecycle:    lifecycle,
			comments:     []Comment{},
			relations:    []Relation{},
		},
	}
	return project
}
