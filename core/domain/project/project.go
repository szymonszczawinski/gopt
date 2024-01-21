package project

import (
	"gopt/core/domain/common/model"
	"time"
)

type Project struct {
	model.TimeTracked
	name        string
	projectKey  string
	description string
	items       []ProjectItem
	owner       projectOwner
	state       projectState
	model.Entity
}

func NewProject(projectKey string, name string, state projectState) Project {
	project := Project{
		Entity:      model.Entity{},
		TimeTracked: model.NewTimeTracked(time.Now(), time.Now()),
		state:       state,
		name:        name,
		projectKey:  projectKey,
		description: "",
		owner: projectOwner{
			id:   1,
			name: "Szymon",
		},
	}
	return project
}

func NewProjectFromRepo(id int, created, updated time.Time, projectKey, name, description string,
	state projectState, items []ProjectItem, owner projectOwner,
) Project {
	project := Project{
		Entity: model.Entity{
			Id: id,
		},
		TimeTracked: model.NewTimeTracked(created, updated),
		state:       state,
		projectKey:  projectKey,
		name:        name,
		description: description,
		items:       items,
		owner:       owner,
	}
	return project
}

func (p Project) GetName() string {
	return p.name
}

func (p Project) GetKey() string {
	return p.projectKey
}

func (p Project) GetDescription() string {
	return p.description
}

func (p Project) GetStateId() int {
	return p.state.id
}

func (p Project) GetLifecycleId() int {
	return p.state.lifecycleId
}

func (p Project) GetOwnerId() int {
	return p.owner.id
}

func (p Project) GetOwner() string {
	return p.owner.name
}

type projectState struct {
	name        string
	id          int
	lifecycleId int
}

func NewProjectState(id, lifecycleId int, name string) projectState {
	return projectState{
		id:          id,
		lifecycleId: lifecycleId,
		name:        name,
	}
}

func (state projectState) String() string {
	return state.name
}

type ProjectItem struct {
	name    string
	itemKey string
	model.Entity
}

func NewProjectItem(id int, name, itemKey string) ProjectItem {
	return ProjectItem{
		name:    name,
		itemKey: itemKey,
		Entity: model.Entity{
			Id: id,
		},
	}
}

type projectOwner struct {
	name string
	id   int
}

func NewProjectOwner(id int, name string) projectOwner {
	return projectOwner{
		id:   id,
		name: name,
	}
}
