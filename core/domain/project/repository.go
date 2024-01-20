package project

import (
	"errors"
	"gopt/coreapi"
	"gopt/coreapi/service"
)

var ErrorCouldNotInsertProject = errors.New("could not insert project")

type IProjectRepository interface {
	service.IComponent
	IProjectQueryRepository
	GetProject(projectId string) coreapi.Result[Project]
	StoreProject(project Project) coreapi.Result[Project]
	UpdateProject(project Project) coreapi.Result[Project]
}

type IProjectQueryRepository interface {
	service.IComponent
	GetProjects() coreapi.Result[[]ProjectListElement]
}
