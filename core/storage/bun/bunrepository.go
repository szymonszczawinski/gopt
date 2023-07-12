package bun

import (
	"context"
	"errors"
	"fmt"
	"gosi/core/storage/dao"
	"gosi/coreapi/service"
	"gosi/issues/domain"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"

	"golang.org/x/sync/errgroup"
)

type disctionaryData struct {
	lifecycleStates map[int]domain.LifecycleState
	lifecycles      map[int]domain.Lifecycle
}

type bunRepository struct {
	lockDb *sync.RWMutex
	db     *bun.DB

	dictionary disctionaryData

	eg  *errgroup.Group
	ctx context.Context
}

func NewRepository(eg *errgroup.Group, ctx context.Context) *bunRepository {
	instance := bunRepository{
		lockDb: &sync.RWMutex{},
		db:     &bun.DB{},
		dictionary: disctionaryData{
			lifecycleStates: map[int]domain.LifecycleState{},
			lifecycles:      map[int]domain.Lifecycle{},
		},
		eg:  eg,
		ctx: ctx,
	}
	return &instance

}
func (self *bunRepository) StartService() {
	if databaseExists() {
		log.Println("Open existing DB")
		self.db, _ = openDatabase(DatabaseDialectSqlite3)
	} else {
		log.Println("Create new DB")
		self.db, _ = createAndInitDb(DatabaseDialectSqlite3, self.ctx)
	}
	self.db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	log.Println("Starting", service.ServiceTypeIRepository)
	self.loadDictionaryData()
}
func (self *bunRepository) Close() {
	self.db.Close()
}

func (self *bunRepository) GetProjects() []domain.Project {
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
func (self bunRepository) GetProject(projectId string) (domain.Project, error) {
	return domain.Project{}, nil
}
func (self *bunRepository) GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error) {
	for _, lc := range self.dictionary.lifecycles {
		if lc.GetName() == string(issueType) {
			return lc, nil
		}
	}
	return domain.Lifecycle{}, errors.New(fmt.Sprintf("Could not find Lifecycle for: %v", string(issueType)))
}
func (self *bunRepository) StoreProject(project domain.Project) (domain.Project, error) {
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

func (self *bunRepository) GetComments() []domain.Comment {
	return nil
}
func (self *bunRepository) StoreComment(comment domain.Comment) (domain.Comment, error) {
	return domain.Comment{}, nil
}

func (self bunRepository) getLifecycle(id int) domain.Lifecycle {
	lifecycle := self.dictionary.lifecycles[id]
	return lifecycle
}

func (self bunRepository) getLifecycleState(id int) domain.LifecycleState {
	lifecyclestate := self.dictionary.lifecycleStates[id]
	return lifecyclestate
}
