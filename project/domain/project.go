package domain

import (
	"gosi/common/model"
	"time"
)

type RelationType string

const (
	RelationTypeCauses     RelationType = "Causes"
	RelationTypeIsCausedBy RelationType = "IsCausedBy"
	RelationTypeIsChildOf  RelationType = "IsChildOf"
	RelationTypeIsParentOf RelationType = "IsParentOf"
)

type Project struct {
	model.Entity
	model.TimeTracked
	model.LivecycleManaged
	Name        string
	ProjectKey  string
	Description string
	Issues      []Issue
}

func NewProject(projectKey string, name string, lifecycle model.Lifecycle) Project {
	project := Project{
		Entity: model.Entity{},
		TimeTracked: model.TimeTracked{
			Created: time.Now(),
			Updated: time.Now(),
		},
		LivecycleManaged: model.LivecycleManaged{
			Lifecycle: lifecycle,
			State:     lifecycle.GetStartState(),
		},
		Name:        name,
		ProjectKey:  projectKey,
		Description: "",
		Issues:      []Issue{},
	}
	return project
}

func NewProjectFromRepo(id int, created time.Time, updated time.Time, projectKey, name, description string,
	state model.LifecycleState, lifecycle model.Lifecycle) Project {
	project := Project{
		Entity: model.Entity{
			Id: id,
		},
		TimeTracked: model.TimeTracked{
			Created: created,
			Updated: updated,
		},
		LivecycleManaged: model.LivecycleManaged{
			Lifecycle: lifecycle,
			State:     state,
		},
		ProjectKey:  projectKey,
		Name:        name,
		Description: description,
	}
	return project
}
