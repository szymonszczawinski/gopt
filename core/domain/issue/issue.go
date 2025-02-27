package issue

import (
	"gopt/core/domain/common"
)

type (
	ParentProject struct {
		id  int
		key string
	}

	Issue struct {
		common.TimeTracked

		issueType   common.IssueType
		project     ParentProject
		itemKey     string
		name        string
		description string
		comments    []Comment
		relations   []Relation
		itemNumber  int
		common.Entity
	}
)

func NewIssueFromRepo(id int, key, name, issueType string, parentId int, parentKey string) Issue {
	issue := Issue{
		Entity: common.Entity{
			Id: id,
		},
		itemKey:   key,
		name:      name,
		issueType: common.IssueType(issueType),
		project: ParentProject{
			id:  parentId,
			key: parentKey,
		},
	}
	return issue
}

func (issue Issue) GetItemKey() string {
	return issue.itemKey
}

func (issue Issue) ParentKey() string {
	return issue.project.key
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

func (issue Issue) GetIssueType() common.IssueType {
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
	issueType common.IssueType
	Issue
}

type Bug struct {
	Issue
	requirement Requirement
}

type Relation struct {
	relationType common.RelationType
	fromIssue    Issue
	toIssue      Issue
	common.Entity
}
