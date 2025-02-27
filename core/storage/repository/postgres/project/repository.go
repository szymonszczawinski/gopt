package project

import (
	"context"
	"errors"
	"gopt/core/domain/common"
	"gopt/core/domain/project"
	"gopt/core/storage/repository/postgres"
	"gopt/coreapi"
	"log"
	"log/slog"
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

	PROJECT_SELECT_BY_KEY = "SELECT p.id, p.created, p.updated, p.name, p.project_key, p.description, p.lifecycle_id, " +
		" lcs.id as state_id, lcs.name as state_name, u.id as owner_id, CONCAT(u.last_name,', ',u.first_name) as owner_name " +
		" from project p " +
		" JOIN lifecyclestate lcs ON lcs.id = p.state_id" +
		" JOIN users u ON u.id = p.created_by_id" +
		" WHERE p.project_key = $1"

	PROJECT_ITEMS_SELECT_BY_PROJECT_KEY = "SELECT i.id, i.created, i.updated, i.name, i.item_key,  " +
		" lcs.name as state_name, lc.name as itemType," +
		" CONCAT(c.last_name,', ',c.first_name) as creator_name, " +
		" CONCAT(a.last_name,', ',a.first_name) as assignee_name " +
		" from issue i " +
		" JOIN lifecyclestate lcs ON lcs.id = i.state_id" +
		" JOIN lifecycle lc ON lc.id = i.lifecycle_id" +
		" JOIN users c ON c.id = i.created_by_id" +
		" JOIN users a ON a.id = i.assigned_to_id" +
		" WHERE i.project_key = $1"
	INSERT_LIFECYCLE_STATE  string = "INSERT INTO lifecyclestate (id, name) VALUES (NULL,?);"
	INSERT_LIFECYCLE        string = "INSERT INTO lifecycle (id, name,start_state_id) VALUES (NULL,?,?);"
	INSERT_STATE_TRANSITION string = "INSERT INTO statetransition (lifecycleid, fromstateid, to_state_id) VALUES (?,?,?);"
	INSERT_PROJECT          string = "INSERT INTO project (id, created, updated, name, project_key, description, state_id, lifecycle_id, created_by_id) " +
		"VALUES (NULL,?,?,?,?,?,?,?,?);"
	INSERT_PROJECT_RETURN_ID string = "INSERT INTO project (created, updated, name, project_key, description, state_id, lifecycle_id, created_by_id) " +
		"VALUES(@created, @updated, @name, @project_key, @description, @state_id, @lifecycle_id, @created_by_id) returning id;"
)

var (
	ErrorProjectNotCreated = errors.New("project not created")
	ErrProjectNotFound     = errors.New("project not found")
)

type projectRepositoryPostgres struct {
	lockDb *sync.RWMutex
	db     postgres.IPostgresDatabase

	eg  *errgroup.Group
	ctx context.Context
}

func NewProjectRepositoryPostgres(eg *errgroup.Group, ctx context.Context, db postgres.IPostgresDatabase) *projectRepositoryPostgres {
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

func (repo projectRepositoryPostgres) GetProject(projectKey string) coreapi.Result[project.Project] {
	result := repo.db.NewSelectOne(PROJECT_SELECT_BY_KEY, projectKey)
	if result == nil {
		return coreapi.NewResult[project.Project](project.Project{}, ErrProjectNotFound)
	}
	// p.id, p.created, p.updated, p.name, p.project_key, p.description, p.lifecycle_id
	var row struct {
		created, updated                                    time.Time
		projectKey, name, ownerName, description, stateName string
		id, lifecycleId, stateId, ownerId                   int
	}

	err := result.Scan(&row.id, &row.created, &row.updated, &row.name, &row.projectKey, &row.description, &row.lifecycleId,
		&row.stateId, &row.stateName, &row.ownerId, &row.ownerName)
	if err != nil {
		// FIXME: remove log
		slog.Info("ERROR:", err, row)
	}
	projectItems := repo.getProjectItems(projectKey)
	p := project.NewProjectFromRepo(row.id, row.created, row.updated, row.projectKey, row.name, row.description,
		project.NewProjectState(row.stateId, row.lifecycleId, row.stateName),
		projectItems.Data(),
		project.NewProjectOwner(row.ownerId, row.ownerName))
	return coreapi.NewResult(p, nil)
}

func (repo projectRepositoryPostgres) StoreProject(p project.Project) coreapi.Result[project.Project] {
	args := pgx.NamedArgs{
		"created":       time.Now(),
		"updated":       time.Now(),
		"name":          p.GetName(),
		"project_key":   p.GetKey(),
		"description":   p.GetDescription(),
		"state_id":      p.GetStateId(),
		"lifecycle_id":  p.GetLifecycleId(),
		"created_by_id": p.GetOwnerId(),
	}
	id, err := repo.db.NewInsertReturninId(INSERT_PROJECT_RETURN_ID, args)
	if err != nil {
		return coreapi.NewResult[project.Project](project.Project{}, errors.Join(ErrorProjectNotCreated, err))
	}
	p.SetId(id)
	return coreapi.NewResult(p, nil)
}

func (repo projectRepositoryPostgres) UpdateProject(p project.Project) coreapi.Result[project.Project] {
	return coreapi.NewResult(project.Project{}, coreapi.ErrorNotImplemented)
}

func (repo projectRepositoryPostgres) getProjectItems(projectKey string) coreapi.Result[[]project.ProjectItem] {
	// FIXME: delete log
	slog.Info("get project items for ", projectKey)
	rows, err := repo.db.NewSelect(PROJECT_ITEMS_SELECT_BY_PROJECT_KEY, projectKey)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return coreapi.NewResult([]project.ProjectItem{}, err)
	}
	defer rows.Close()
	projectItems := []project.ProjectItem{}
	var row struct {
		created, updated                                  time.Time
		itemKey, name, state, creator, assingee, itemType string
		id                                                int
	}

	pgx.ForEachRow(rows, []any{&row.id, &row.created, &row.updated, &row.name, &row.itemKey, &row.state, &row.itemType, &row.creator, &row.assingee}, func() error {
		projectItems = append(projectItems, project.NewProjectItem(row.id, row.itemKey, row.name, common.IssueType(row.itemType),
			row.state, row.created, row.updated, row.creator, row.assingee))
		return nil
	})
	return coreapi.NewResult(projectItems, nil)
}
