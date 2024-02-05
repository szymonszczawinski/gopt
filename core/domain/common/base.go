package common

import "time"

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
	IssueTypeTask        IssueType = "Task"
)

type Named interface {
	GetName() string
}

type Entity struct {
	Id int
}

func (e Entity) GetId() int {
	return e.Id
}

func (e *Entity) SetId(id int) {
	e.Id = id
}

type TimeTracked struct {
	created time.Time
	updated time.Time
}

func NewTimeTracked(created, updated time.Time) TimeTracked {
	return TimeTracked{
		created: created,
		updated: updated,
	}
}

func (tt TimeTracked) GetCreationTime() time.Time {
	return tt.created
}

func (tt TimeTracked) GetLastUpdateTime() time.Time {
	return tt.updated
}

func (tt *TimeTracked) SetLastUpdateTime(lastUpdated time.Time) {
	tt.updated = lastUpdated
}
