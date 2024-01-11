package postgres

import (
	"context"
	"gosi/core/domain/project"
	"gosi/coreapi"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"golang.org/x/sync/errgroup"
)

const (
	PROJECTS_SELECT_ALL = "SELECT project_row.id, project_row.project_key, project_row.name, " +
		" project_row.created, project_row.updated," +
		" lifecyclestate.name as state_name, CONCAT(users.last_name,', ',users.first_name) as owner_name " +
		" FROM project AS project_row " +
		" JOIN lifecyclestate ON lifecyclestate.id = project_row.state_id" +
		" JOIN users ON users.id = project_row.created_by_id"
)

type projectRepositoryPostgres struct {
	lockDb *sync.RWMutex
	db     IPostgresDatabase

	eg  *errgroup.Group
	ctx context.Context
}

func NewProjectRepositoryPostgres(eg *errgroup.Group, ctx context.Context, db IPostgresDatabase) project.IProjectRepository {
	instance := projectRepositoryPostgres{
		lockDb: &sync.RWMutex{},
		db:     db,
		eg:     eg,
		ctx:    ctx,
	}
	return &instance
}

func (repo *projectRepositoryPostgres) StartComponent() {
}

func (repo projectRepositoryPostgres) GetProjects() coreapi.Result[[]project.ProjectListElement] {
	rows, err := repo.db.NewSelect(PROJECTS_SELECT_ALL)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return coreapi.NewResult([]project.ProjectListElement{}, err)
	}
	defer rows.Close()
	projects := []project.ProjectListElement{}
	var row struct {
		created, updated               time.Time
		projectKey, name, state, owner string
		id                             int
	}

	pgx.ForEachRow(rows, []any{&row.id, &row.projectKey, &row.name, &row.created, &row.updated, &row.state, &row.owner}, func() error {
		projects = append(projects, project.NewProjectListElement(row.id, row.projectKey, row.name, row.state, row.owner, row.created, row.updated))
		return nil
	})

	return coreapi.NewResult(projects, nil)
}

func (repo projectRepositoryPostgres) GetProject(projectId string) coreapi.Result[project.Project] {
	return coreapi.NewResult(project.Project{}, coreapi.ErrorNotImplemented)
}

func (repo projectRepositoryPostgres) StoreProject(p project.Project) coreapi.Result[project.Project] {
	return coreapi.NewResult(project.Project{}, coreapi.ErrorNotImplemented)
}

func (repo projectRepositoryPostgres) UpdateProject(p project.Project) coreapi.Result[project.Project] {
	return coreapi.NewResult(project.Project{}, coreapi.ErrorNotImplemented)
}
