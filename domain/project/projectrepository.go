package project

import (
	"context"
	"errors"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"gosi/domain/common/model"
	"gosi/domain/project/dao"
	"gosi/domain/project/domain"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

type IProjectRepository interface {
	service.IComponent
	GetProjects() []domain.Project
	GetProject(projectId string) (domain.Project, error)
	GetLifecycle() (model.Lifecycle, error)
	StoreProject(project domain.Project) (domain.Project, error)
	UpdateProject(project domain.Project) (domain.Project, error)
}

type disctionaryData struct {
	lifecycleStates map[int]model.LifecycleState
	lifecycles      map[int]model.Lifecycle
}

type projectRepository struct {
	lockDb *sync.RWMutex
	db     storage.IBunDatabase

	dictionary disctionaryData

	eg  *errgroup.Group
	ctx context.Context
}

func NewProjectRepository(eg *errgroup.Group, ctx context.Context, db storage.IBunDatabase) *projectRepository {
	instance := projectRepository{
		lockDb: &sync.RWMutex{},
		db:     db,
		dictionary: disctionaryData{
			lifecycleStates: map[int]model.LifecycleState{},
			lifecycles:      map[int]model.Lifecycle{},
		},
		eg:  eg,
		ctx: ctx,
	}
	return &instance
}

func (self *projectRepository) StartComponent() {
	self.loadDictionaryData()
}

func (self *projectRepository) GetProjects() []domain.Project {
	var (
		projectsRows []dao.ProjectRow
		projects     []domain.Project = []domain.Project{}
	)
	self.lockDb.RLock()
	defer self.lockDb.RUnlock()

	err := self.db.NewSelect().Model(&projectsRows).Scan(self.ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range projectsRows {
		projects = append(projects, domain.NewProjectFromRepo(row.Id, row.Created, row.Updated, row.ProjectKey, row.Name,
			row.Description, self.getLifecycleState(row.StateId), self.getLifecycle(row.LifecycleId)))
		log.Println(projects)
	}
	return projects

}

func (self projectRepository) GetProject(projectId string) (domain.Project, error) {
	self.lockDb.RLock()
	self.lockDb.RUnlock()

	return domain.Project{}, nil
}

func (self *projectRepository) GetLifecycle() (model.Lifecycle, error) {
	for _, lc := range self.dictionary.lifecycles {
		if lc.GetName() == "Project" {
			return lc, nil
		}
	}
	return model.Lifecycle{}, errors.New("Could not find Project Lifecycle")
}

func (self *projectRepository) StoreProject(project domain.Project) (domain.Project, error) {
	self.lockDb.Lock()
	self.lockDb.Unlock()
	dao := &dao.ProjectRow{
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
		return domain.Project{}, err
	}
	id, err := res.LastInsertId()
	log.Println("RES :: ", id, " :: ", err)
	return domain.Project{}, nil
}

func (self *projectRepository) UpdateProject(p domain.Project) (domain.Project, error) {
	log.Fatal("Not Implemented")
	return domain.Project{}, nil
}

func (self projectRepository) getLifecycle(id int) model.Lifecycle {
	lifecycle := self.dictionary.lifecycles[id]
	return lifecycle
}

func (self projectRepository) getLifecycleState(id int) model.LifecycleState {
	lifecyclestate := self.dictionary.lifecycleStates[id]
	return lifecyclestate
}
func (self *projectRepository) loadDictionaryData() {
	self.loadLifecycles()
}

func (self *projectRepository) loadLifecycles() {
	var (
		lifecycleStatesRows []dao.LifecycleStateRow
		lifecyclesRows      []dao.LifecycleRow
	)
	err := self.db.NewSelect().Model(&lifecycleStatesRows).Scan(self.ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range lifecycleStatesRows {
		self.dictionary.lifecycleStates[row.Id] = model.NewLifecycleState(row.Id, row.Name)
	}

	err = self.db.NewSelect().Model(&lifecyclesRows).Scan(self.ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range lifecyclesRows {
		self.dictionary.lifecycles[row.Id] = model.NewLifeCycleBuilder(row.Id, row.Name, self.dictionary.lifecycleStates[row.StartStateId]).Build()
	}
}
