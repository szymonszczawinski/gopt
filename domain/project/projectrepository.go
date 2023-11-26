package project

import (
	"errors"
	"gosi/coreapi"
	"gosi/coreapi/service"
)

const (
	PROJECTS_SELECT_ALL = "SELECT project_row.id, project_row.project_key, project_row.name, " +
		" project_row.created, project_row.updated," +
		" lifecyclestate.name as state_name, CONCAT(users.last_name,', ',users.first_name) as owner_name " +
		" FROM project AS project_row " +
		" JOIN lifecyclestate ON lifecyclestate.id = project_row.state_id" +
		" JOIN users ON users.id = project_row.owner_id"
)

var ErrorCouldNotInsertProject = errors.New("could not insert project")

type IProjectRepository interface {
	service.IComponent
	GetProject(projectId string) coreapi.Result[Project]
	StoreProject(project Project) coreapi.Result[Project]
	UpdateProject(project Project) coreapi.Result[Project]
}

type IProjectQueryRepository interface {
	service.IComponent
	GetProjects() coreapi.Result[[]ProjectListElement]
}
