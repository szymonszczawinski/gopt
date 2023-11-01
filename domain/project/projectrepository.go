package project

import (
	"context"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

type IProjectRepository interface {
	service.IComponent
	GetProjects() []Project
	GetProject(projectId string) (Project, error)
	StoreProject(project Project) (Project, error)
	UpdateProject(project Project) (Project, error)
	GetProjectState() (ProjectState, error)
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
	projectState, _ := repo.GetProjectState()
	for _, row := range projectsRows {
		projects = append(projects, NewProjectFromRepo(row.Id, row.Created, row.Updated, row.ProjectKey, row.Name,
			row.Description, projectState))
		log.Println(projects)
	}
	return projects

}

func (repo projectRepository) GetProject(projectId string) (Project, error) {
	repo.lockDb.RLock()
	defer repo.lockDb.RUnlock()

	return Project{}, nil
}

func (repo *projectRepository) StoreProject(project Project) (Project, error) {
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
		return Project{}, err
	}
	id, err := res.LastInsertId()
	log.Println("RES :: ", id, " :: ", err)
	return Project{}, nil
}

func (repo *projectRepository) UpdateProject(p Project) (Project, error) {
	log.Fatal("Not Implemented")
	return Project{}, nil
}

func (repo *projectRepository) GetProjectState() (ProjectState, error) {
	//TODO: to implement
	return NewProjectState(1, 1, "Open"), nil
}
