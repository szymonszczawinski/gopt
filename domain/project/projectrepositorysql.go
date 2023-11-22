package project

import (
	"context"
	"gosi/core/storage/sql"
	"gosi/coreapi"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"golang.org/x/sync/errgroup"
)

type projectRepositorySql struct {
	lockDb *sync.RWMutex
	db     sql.IPostgresDatabase

	eg  *errgroup.Group
	ctx context.Context
}

func NewProjectRepositorySql(eg *errgroup.Group, ctx context.Context, db sql.IPostgresDatabase) *projectRepositorySql {
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

func (repo projectRepositorySql) GetProjects() coreapi.Result[[]ProjectListElement] {
	rows, err := repo.db.NewSelect(PROJECTS_SELECT_ALL)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return coreapi.NewResult([]ProjectListElement{}, err)
	}
	defer rows.Close()
	projects := []ProjectListElement{}
	var row struct {
		id                             int
		projectKey, name, state, owner string
		created, updated               time.Time
	}

	pgx.ForEachRow(rows, []any{&row.id, &row.projectKey, &row.name, &row.created, &row.updated, &row.state, &row.owner}, func() error {
		projects = append(projects, NewProjectListElement(row.id, row.projectKey, row.name, row.state, row.owner, row.created, row.updated))
		return nil
	})

	return coreapi.NewResult(projects, nil)
}

func (repo projectRepositorySql) GetProject(projectId string) coreapi.Result[Project] {
	return coreapi.NewResult(Project{}, coreapi.ErrorNotImplemented)
}

func (repo projectRepositorySql) StoreProject(project Project) coreapi.Result[Project] {
	return coreapi.NewResult(Project{}, coreapi.ErrorNotImplemented)
}

func (repo projectRepositorySql) UpdateProject(project Project) coreapi.Result[Project] {
	return coreapi.NewResult(Project{}, coreapi.ErrorNotImplemented)
}
