package project

import (
	"context"
	"errors"
	"gosi/coreapi"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	PROJECTS_SELECT_ALL = "SELECT project_row.*, lifecyclestate.name as state_name, CONCAT(users.last_name,', ',users.first_name) as owner_name " +
		" FROM project AS project_row " +
		" JOIN lifecyclestate ON lifecyclestate.id = project_row.state_id" +
		" JOIN users ON users.id = project_row.owner_id"
)

var ErrorCouldNotInsertProject = errors.New("could not insert project")

type IProjectRepository interface {
	service.IComponent
	GetProjects() coreapi.Result[[]ProjectListElement]
	GetProject(projectId string) coreapi.Result[Project]
	StoreProject(project Project) coreapi.Result[Project]
	UpdateProject(project Project) coreapi.Result[Project]
}

type projectRepository struct {
	lockDb *sync.RWMutex
	db     storage.IBunDatabase

	eg  *errgroup.Group
	ctx context.Context
}

func NewProjectRepository(eg *errgroup.Group, ctx context.Context, db storage.IBunDatabase) *projectRepository {
	instance := projectRepository{
		lockDb: &sync.RWMutex{},
		db:     db,
		eg:     eg,
		ctx:    ctx,
	}
	return &instance
}

func (repo *projectRepository) StartComponent() {
}

func (repo *projectRepository) GetProjects() coreapi.Result[[]ProjectListElement] {
	var projectsRows []ProjectRow
	var projects []ProjectListElement

	err := repo.db.NewRaw(PROJECTS_SELECT_ALL).Scan(repo.ctx, &projectsRows)
	if err != nil {
		log.Println(err)
		return coreapi.NewResult([]ProjectListElement{}, err)
	}
	for _, row := range projectsRows {
		projects = append(projects,
			NewProjectListElement(row.Id, row.ProjectKey, row.Name,
				row.StateName, row.OwnerName, row.Created, row.Updated))
	}
	return coreapi.NewResult(projects, nil)
}

func (repo projectRepository) GetProject(projectId string) coreapi.Result[Project] {
	// TODO: to implement
	return coreapi.NewResult[Project](Project{}, coreapi.ErrorNotImplemented)
}

func (repo *projectRepository) StoreProject(project Project) coreapi.Result[Project] {
	repo.lockDb.Lock()
	defer repo.lockDb.Unlock()
	dao := &ProjectRow{
		Name:        project.Name,
		ProjectKey:  project.ProjectKey,
		Description: project.Description,
		StateId:     project.State.id,
		LifecycleId: project.State.lifecycleId,
		CreatedById: 0,
		OwnerId:     0,
	}
	res, err := repo.db.NewInsert().Model(dao).Returning("id").Exec(repo.ctx)
	if err != nil {
		log.Println(errors.Join(ErrorCouldNotInsertProject, err))
		return coreapi.NewResult[Project](Project{}, errors.Join(ErrorCouldNotInsertProject, err))
	}
	id, err := res.LastInsertId()
	log.Println("RES :: ", id, " :: ", err)
	return coreapi.NewResult[Project](Project{}, nil)
}

func (repo *projectRepository) UpdateProject(p Project) coreapi.Result[Project] {
	// TODO: to omplement UpdateProject
	return coreapi.NewResult[Project](Project{}, coreapi.ErrorNotImplemented)
}

func (repo *projectRepository) getProjectState() coreapi.Result[ProjectState] {
	// TODO: to implement GetProjectState
	return coreapi.NewResult[ProjectState](NewProjectState(1, 1, "Open"), nil)
}
