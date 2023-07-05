package bun

import (
	"context"
	"database/sql"
	"gosi/core/storage/dao"
	"gosi/coreapi/service"
	"gosi/issues/domain"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
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
	instance := new(bunRepository)
	instance.ctx = ctx
	instance.eg = eg
	instance.dictionary = disctionaryData{
		lifecycleStates: map[int]domain.LifecycleState{},
		lifecycles:      map[int]domain.Lifecycle{},
	}
	instance.lockDb = &sync.RWMutex{}
	return instance

}
func (self *bunRepository) StartService() {
	dbfile := os.Getenv("DATABASE_FILE_NAME")
	db, errOpenDB := sql.Open("sqlite3", dbfile)
	if errOpenDB != nil {
		log.Fatal(errOpenDB.Error())
	} else {
		self.db = bun.NewDB(db, sqlitedialect.New())
	}
	log.Println("Starting", service.ServiceTypeIRepository)
	self.loadDictionaryData()
}
func (self bunRepository) Close() {
	self.db.Close()
}

func (self *bunRepository) GetProjects() []domain.Project {
	self.lockDb.RLock()
	defer self.lockDb.RUnlock()
	rows, err := self.db.NewSelect().
		ColumnExpr("id").ColumnExpr("created").ColumnExpr("updated").ColumnExpr("name").
		ColumnExpr("itemkey").ColumnExpr("itemnumber").ColumnExpr("description").
		ColumnExpr("stateid").ColumnExpr("lifecycleid").ColumnExpr("createdbyid").
		TableExpr("project").
		Rows(self.ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	projects := []domain.Project{}

	for rows.Next() {
		var row dao.ProjectRow
		err = rows.Scan(&row.Id, &row.Created, &row.Updated, &row.Name, &row.ItemKey, &row.ItemNumber,
			&row.Description, &row.StateId, &row.LifecycleId, &row.CreatedById)
		if err != nil {
			log.Println(err.Error())
		} else {
			projects = append(projects, domain.NewProjectFromRepo(row.Id, row.Created, row.Updated, row.ItemKey, row.ItemNumber, row.Name,
				row.Description, self.getLifecycleState(row.StateId), self.getLifecycle(row.LifecycleId)))
		}
	}
	return projects

}
func (self bunRepository) GetProject(projectId string) (domain.Project, error) {
	return domain.Project{}, nil
}
func (self *bunRepository) GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error) {
	return domain.Lifecycle{}, nil
}
func (self *bunRepository) StoreProject(project domain.Project) (domain.Project, error) {
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

func (self *bunRepository) loadDictionaryData() {
	self.loadLifecycles()
}

func (self *bunRepository) loadLifecycles() {
	rows, err := self.db.NewSelect().
		ColumnExpr("id").ColumnExpr("name").
		TableExpr("lifecyclestate").
		Rows(self.ctx)

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println("ERROR::", err.Error())
		} else {
			self.dictionary.lifecycleStates[id] = domain.NewLifecycleState(id, name)
		}
	}
	rows.Close()
	rows, err = self.db.NewSelect().
		ColumnExpr("id").ColumnExpr("name").ColumnExpr("startstateid").
		TableExpr("lifecycle").
		Rows(self.ctx)
	log.Println("LIFECYCLE STATES LOADED: ", len(self.dictionary.lifecycleStates))
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var startStateId int
		rows.Scan(&id, &name, &startStateId)
		if err != nil {
			log.Println("ERROR::", err.Error())
		} else {
			self.dictionary.lifecycles[id] = domain.NewLifeCycleBuilder(id, name, self.dictionary.lifecycleStates[startStateId]).Build()
		}
	}
	rows.Close()
	log.Println("LIFECYCLES LOADED: ", len(self.dictionary.lifecycles))
}
