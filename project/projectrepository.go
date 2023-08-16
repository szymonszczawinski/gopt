package project

import (
	"context"
	"errors"
	"fmt"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"gosi/project/dao"
	"gosi/project/domain"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

type IProjectRepository interface {
	service.IComponent
	GetProjects() []domain.Project
	GetProject(projectId string) (domain.Project, error)
	GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error)
	StoreProject(project domain.Project) (domain.Project, error)
	GetComments() []domain.Comment
	StoreComment(comment domain.Comment) (domain.Comment, error)
}

type disctionaryData struct {
	lifecycleStates map[int]domain.LifecycleState
	lifecycles      map[int]domain.Lifecycle
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
			lifecycleStates: map[int]domain.LifecycleState{},
			lifecycles:      map[int]domain.Lifecycle{},
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
		projects = append(projects, domain.NewProjectFromRepo(row.Id, row.Created, row.Updated, row.ItemKey, row.ItemNumber, row.Name,
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
func (self *projectRepository) GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error) {
	for _, lc := range self.dictionary.lifecycles {
		if lc.GetName() == string(issueType) {
			return lc, nil
		}
	}
	return domain.Lifecycle{}, errors.New(fmt.Sprintf("Could not find Lifecycle for: %v", string(issueType)))
}
func (self *projectRepository) StoreProject(project domain.Project) (domain.Project, error) {
	self.lockDb.Lock()
	self.lockDb.Unlock()
	dao := &dao.ProjectRow{
		Name:        project.GetName(),
		ItemKey:     project.GetItemKey(),
		ItemNumber:  project.GetItemNumber(),
		Description: project.GetDescription(),
		StateId:     project.GetState().GetId(),
		LifecycleId: project.GetLifecycle().GetId(),
		CreatedById: 0,
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

func (self *projectRepository) GetComments() []domain.Comment {
	self.lockDb.RLock()
	self.lockDb.RUnlock()

	return nil
}
func (self *projectRepository) StoreComment(comment domain.Comment) (domain.Comment, error) {
	self.lockDb.Lock()
	self.lockDb.Unlock()

	return domain.Comment{}, nil
}

func (self projectRepository) getLifecycle(id int) domain.Lifecycle {
	lifecycle := self.dictionary.lifecycles[id]
	return lifecycle
}

func (self projectRepository) getLifecycleState(id int) domain.LifecycleState {
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
		self.dictionary.lifecycleStates[row.Id] = domain.NewLifecycleState(row.Id, row.Name)
	}

	err = self.db.NewSelect().Model(&lifecyclesRows).Scan(self.ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range lifecyclesRows {
		self.dictionary.lifecycles[row.Id] = domain.NewLifeCycleBuilder(row.Id, row.Name, self.dictionary.lifecycleStates[row.StartStateId]).Build()
	}
}
