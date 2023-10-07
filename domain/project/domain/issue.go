package domain

import "gosi/domain/common/model"

type IssueType string

const (
	IssueTypeRequirement IssueType = "Requirement"
	IssueTypeBug         IssueType = "Bug"
)

type Issue struct {
	model.Entity
	model.TimeTracked
	model.LivecycleManaged

	itemKey     string
	itemNumber  int
	name        string
	description string
	project     Project
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
	model.Entity
	relationType RelationType
	fromIssue    Issue
	toIssue      Issue
}
