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

func (self *projectRepository) StartComponent() {
}

func (self *projectRepository) GetProjects() []Project {
	var (
		projectsRows []ProjectRow
		projects     []Project = []Project{}
	)
	self.lockDb.RLock()
	defer self.lockDb.RUnlock()

	err := self.db.NewSelect().Model(&projectsRows).Scan(self.ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range projectsRows {
		projects = append(projects, NewProjectFromRepo(row.Id, row.Created, row.Updated, row.ProjectKey, row.Name,
			row.Description, self.getLifecycleState(row.StateId), self.getLifecycle(row.LifecycleId)))
		log.Println(projects)
	}
	return projects

}

func (self projectRepository) GetProject(projectId string) (Project, error) {
	self.lockDb.RLock()
	self.lockDb.RUnlock()

	return Project{}, nil
}

func (self *projectRepository) StoreProject(project Project) (Project, error) {
	self.lockDb.Lock()
	self.lockDb.Unlock()
	dao := &ProjectRow{
		Name:        project.Name,
		ProjectKey:  project.ProjectKey,
		Description: project.Description,
		StateId:     project.GetState().GetId(),
		LifecycleId: project.GetLifecycle().GetId(),
		CreatedById: 0,
		OwnerId:     0,
	}
	res, err := self.db.NewInsert().Model(dao).Returning("id").Exec(self.ctx)
	if err != nil {
		log.Println("ERROR when insert project", err.Error())
		return Project{}, err
	}
	id, err := res.LastInsertId()
	log.Println("RES :: ", id, " :: ", err)
	return Project{}, nil
}

func (self *projectRepository) UpdateProject(p Project) (Project, error) {
	log.Fatal("Not Implemented")
	return Project{}, nil
}
