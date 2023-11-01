package domain

import "gosi/domain/common/model"

type IssueType string

type RelationType string

const (
	RelationTypeCauses     RelationType = "Causes"
	RelationTypeIsCausedBy RelationType = "IsCausedBy"
	RelationTypeIsChildOf  RelationType = "IsChildOf"
	RelationTypeIsParentOf RelationType = "IsParentOf"
)
const (
	IssueTypeRequirement IssueType = "Requirement"
	IssueTypeBug         IssueType = "Bug"
)

type ParentProject struct {
}

type Issue struct {
	model.Entity
	model.TimeTracked

	itemKey     string
	itemNumber  int
	name        string
	description string
	project     ParentProject
	issueType   IssueType
	comments    []Comment
	relations   []Relation
}

func (issue Issue) GetItemKey() string {
	return issue.itemKey
}
func (issue Issue) GetItemNumber() int {
	return issue.itemNumber
}

func (issue Issue) GetName() string {
	return issue.name
}

func (issue Issue) GetDescription() string {
	return issue.description
}

func (issue Issue) GetIssueType() IssueType {
	return issue.issueType
}

func (issue *Issue) GetAndIncrementItemNumber() int {
	issue.itemNumber += 1
	return issue.itemNumber
}

func (issue *Issue) AddComment(comment Comment) {
	issue.comments = append(issue.comments, comment)
}

type Requirement struct {
	Issue
	project   ParentProject
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
