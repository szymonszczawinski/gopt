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

func NewProjectFromRepo(id int, created time.Time, updated time.Time, projectKey, name, description string,
	state projectState, items []ProjectItem,
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
	Name    string
	ItemKey string
	model.Entity
}

type projectOwner struct {
	id   int
	name string
}
