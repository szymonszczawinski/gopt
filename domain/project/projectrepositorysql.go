package project

import (
	"context"
	"gosi/core/storage/sql"
	"gosi/coreapi"
	"log"
	"sync"
	"time"

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
	var row ProjectListElement
	var created, updated time.Time
	for rows.Next() {
		err := rows.Scan(&row.Id, &row.ProjectKey, &row.Name, &created, &updated, &row.State, &row.Owner)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(row)
	}
	// pgx.ForEachRow(rows, []any{&row.Id, &row.ProjectKey, &row.Name, &row.Created, &row.Updated, &row.State, &row.Owner}, func() error {
	// 	log.Printf("====== ROW: %v\v", row)
	// 	projects = append(projects, row)
	// 	return nil
	// })

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
