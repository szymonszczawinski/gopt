package project

import (
	"gosi/domain/common/model"
	"time"
)

type Project struct {
	model.Entity
	model.TimeTracked
	Name        string
	ProjectKey  string
	Description string
	State       ProjectState
	Items       []ProjectItem
}

func NewProject(projectKey string, name string, state ProjectState) Project {
	project := Project{
		Entity: model.Entity{},
		TimeTracked: model.TimeTracked{
			Created: time.Now(),
			Updated: time.Now(),
		},
		State:       state,
		Name:        name,
		ProjectKey:  projectKey,
		Description: "",
	}
	return project
}

func NewProjectFromRepo(id int, created time.Time, updated time.Time, projectKey, name, description string,
	state ProjectState, items []ProjectItem,
) Project {
	project := Project{
		Entity: model.Entity{
			Id: id,
		},
		TimeTracked: model.TimeTracked{
			Created: created,
			Updated: updated,
		},
		State:       state,
		ProjectKey:  projectKey,
		Name:        name,
		Description: description,
		Items:       items,
	}
	return project
}

type ProjectState struct {
	id          int
	lifecycleId int
	name        string
}

func (state ProjectState) String() string {
	return state.name
}

func NewProjectState(id, lifecycleId int, name string) ProjectState {
	return ProjectState{
		id:          id,
		lifecycleId: lifecycleId,
		name:        name,
	}
}

type ProjectItem struct {
	model.Entity
	Name    string
	ItemKey string
}
