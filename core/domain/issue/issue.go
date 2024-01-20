package domain

import "gopt/core/domain/common/model"

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

type (
	ParentProject struct{}

	Issue struct {
		model.TimeTracked

		issueType   IssueType
		project     ParentProject
		itemKey     string
		name        string
		description string
		comments    []Comment
		relations   []Relation
		itemNumber  int
		model.Entity
	}
)

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
	project   ParentProject
	issueType IssueType
	Issue
}

type Bug struct {
	Issue
	requirement Requirement
}

type Relation struct {
	relationType RelationType
	fromIssue    Issue
	toIssue      Issue
	model.Entity
}
