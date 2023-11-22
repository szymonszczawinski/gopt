package project

import (
	"context"
	"gosi/coreapi"
	"gosi/coreapi/storage"
	"sync"

	"golang.org/x/sync/errgroup"
)

type projectRepositorySql struct {
	lockDb *sync.RWMutex
	db     storage.ISqlDatabase

	eg  *errgroup.Group
	ctx context.Context
}

func NewProjectRepositorySql(eg *errgroup.Group, ctx context.Context, db storage.ISqlDatabase) *projectRepositorySql {
	instance := projectRepositorySql{
		lockDb: &sync.RWMutex{},
		db:     db,
		eg:     eg,
		ctx:    ctx,
	}
	return &instance
}

func (repo *projectRepositorySql) StartComponent() {
}

func GetProjects() coreapi.Result[[]ProjectListElement] {
	return coreapi.NewResult([]ProjectListElement{}, coreapi.ErrorNotImplemented)
}

func GetProject(projectId string) coreapi.Result[Project] {
	return coreapi.NewResult(Project{}, coreapi.ErrorNotImplemented)
}

func StoreProject(project Project) coreapi.Result[Project] {
	return coreapi.NewResult(Project{}, coreapi.ErrorNotImplemented)
}

func UpdateProject(project Project) coreapi.Result[Project] {
	return coreapi.NewResult(Project{}, coreapi.ErrorNotImplemented)
}
