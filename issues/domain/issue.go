package domain

import (
	"time"
)

type IssueType string
type RelationType string

const (
	IssueTypeProject     IssueType = "Project"
	IssueTypeRequirement IssueType = "Requirement"
	IssueTypeBug         IssueType = "Bug"
)

const (
	RelationTypeCauses     RelationType = "Causes"
	RelationTypeIsCausedBy RelationType = "IsCausedBy"
	RelationTypeIsChildOf  RelationType = "IsChildOf"
	RelationTypeIsParentOf RelationType = "IsParentOf"
)

type Issue struct {
	Entity
	TimeTracked
	LivecycleManaged
	itemKey     string
	itemNumber  int
	name        string
	description string
	issueType   IssueType
	comments    []Comment
	relations   []Relation
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

func (self *Issue) GetAndIncrementItemNumber() int {
	self.itemNumber += 1
	return self.itemNumber
}

func (self *Issue) AddComment(comment Comment) {
	self.comments = append(self.comments, comment)
}

type Project struct {
	Issue
	requirements []Requirement
}

func NewProject(projectKey string, name string, lifecycle Lifecycle) Project {
	project := Project{
		Issue: Issue{
			Entity: Entity{},
			TimeTracked: TimeTracked{
				created: time.Now(),
				updated: time.Now(),
			},
			LivecycleManaged: LivecycleManaged{
				lifecycle: lifecycle,
				state:     lifecycle.startState,
			},
			itemKey:     projectKey,
			itemNumber:  1,
			name:        name,
			description: "",
			issueType:   IssueTypeProject,
			comments:    []Comment{},
		},
		requirements: []Requirement{},
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
			LivecycleManaged: LivecycleManaged{
				lifecycle: lifecycle,
				state:     state,
			},
			itemKey:     itemKey,
			itemNumber:  itemNumber,
			name:        name,
			description: description,
			issueType:   IssueTypeProject,
			comments:    []Comment{},
			relations:   []Relation{},
		},
	}
	return project
}

// func (self Project) String() string {
// 	return fmt.Sprintf("Project[id:%v; key:%v; name:%v; state:%v\n comments:%v]",
// 		self.GetId(), self.GetItemKey(), self.GetName(), self.currentState, self.comments)
// }

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
