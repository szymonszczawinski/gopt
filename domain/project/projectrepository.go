package project

import (
	"context"
	"gosi/coreapi"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

type IProjectRepository interface {
	service.IComponent
	GetProjects() []Project
	GetProject(projectId string) coreapi.Result[Project]
	StoreProject(project Project) coreapi.Result[Project]
	UpdateProject(project Project) coreapi.Result[Project]
	GetProjectState() coreapi.Result[ProjectState]
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

func (repo *projectRepository) GetProjects() []Project {
	var (
		projectsRows []ProjectRow
		projects     []Project = []Project{}
	)
	repo.lockDb.RLock()
	defer repo.lockDb.RUnlock()

	err := repo.db.NewSelect().Model(&projectsRows).Scan(repo.ctx)
	if err != nil {
		log.Fatal(err)
	}
	//TODO: tio implement
	resultState := repo.GetProjectState()
	for _, row := range projectsRows {
		projects = append(projects, NewProjectFromRepo(row.Id, row.Created, row.Updated, row.ProjectKey, row.Name,
			row.Description, resultState.Data()))
		log.Println(projects)
	}
	return projects

}

func (repo projectRepository) GetProject(projectId string) coreapi.Result[Project] {
	repo.lockDb.RLock()
	defer repo.lockDb.RUnlock()

	return coreapi.NewResult[Project](Project{}, nil)
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
		log.Println("ERROR when insert project", err.Error())
		return coreapi.NewResult[Project](Project{}, err)
	}
	id, err := res.LastInsertId()
	log.Println("RES :: ", id, " :: ", err)
	return coreapi.NewResult[Project](Project{}, nil)
}

func (repo *projectRepository) UpdateProject(p Project) coreapi.Result[Project] {
	log.Fatal("Not Implemented")
	return coreapi.NewResult[Project](Project{}, nil)
}

func (repo *projectRepository) GetProjectState() coreapi.Result[ProjectState] {
	//TODO: to implement
	return coreapi.NewResult[ProjectState](NewProjectState(1, 1, "Open"), nil)
}
