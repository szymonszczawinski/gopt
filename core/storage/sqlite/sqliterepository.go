package sqlite

import (
	"database/sql"
	"gosi/core/storage/sql/query"
	"gosi/coreapi/service"
	"gosi/issues/domain"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"gosi/issues/dao"
)

type sqliteRepository struct {
	database        *sql.DB
	mux             *sync.RWMutex
	lifecycleStates map[int]domain.LifecycleState
	lifecycles      map[int]domain.Lifecycle
}

func NewSqliteRepository() *sqliteRepository {
	instance := sqliteRepository{
		lifecycleStates: make(map[int]domain.LifecycleState),
		lifecycles:      make(map[int]domain.Lifecycle),
	}

	return &instance

}

func (self *sqliteRepository) StartService() {
	log.Println("Starting", service.ComponentTypeSqlite3Repository)
	dbfile := os.Getenv("DATABASE_FILE_NAME")
	db, errOpenDB := sql.Open("sqlite3", dbfile)
	if errOpenDB != nil {
		log.Println(errOpenDB.Error())
	} else {
		self.database = db
		self.mux = &sync.RWMutex{}
		// initaliseDatabase(instance.database)
	}
	self.loadDictionaryData()

}

func (self sqliteRepository) Close() {
	self.database.Close()
}

func (self *sqliteRepository) GetProjects() []domain.Project {
	self.mux.RLock()
	defer self.mux.RUnlock()
	rows, err := self.database.Query(query.PROJECT_SELECT_ALL)
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
		}
		log.Println("LIFECYCLE: ", self.lifecycles[row.LifecycleId])
		log.Println("LIFECYCLE STATE: ", self.lifecycleStates[row.StateId])
		projects = append(projects, domain.NewProjectFromRepo(row.Id, row.Created, row.Updated, row.ItemKey, row.ItemNumber, row.Name,
			row.Description, self.lifecycleStates[row.StateId], self.lifecycles[row.LifecycleId]))
	}
	log.Println(projects)
	return projects
}
func (self *sqliteRepository) GetProject(projectId string) (domain.Project, error) {
	return domain.Project{}, nil
}

func (self *sqliteRepository) GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error) {
	return domain.Lifecycle{}, nil
}

func (self *sqliteRepository) StoreProject(project domain.Project) (domain.Project, error) {
	return domain.Project{}, nil
}
func (self *sqliteRepository) GetComments() []domain.Comment {
	return nil
}
func (self *sqliteRepository) StoreComment(comment domain.Comment) (domain.Comment, error) {
	return domain.Comment{}, nil
}

func (self *sqliteRepository) loadDictionaryData() {
	self.loadLifecycles()
}

func (self *sqliteRepository) loadLifecycles() {
	rows, err := self.database.Query(query.LIFECYCLE_STATE_SELECT_ALL)

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Println("ERROR::", err.Error())
		}
		self.lifecycleStates[id] = domain.NewLifecycleState(id, name)
	}
	rows.Close()
	log.Println("LIFECYCLE STATES LOADED: ", len(self.lifecycleStates))
	rows, err = self.database.Query(query.LIFECYCLE_SELECT_ALL)
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
		}

		self.lifecycles[id] = domain.NewLifeCycleBuilder(id, name, self.lifecycleStates[startStateId]).Build()
	}
	rows.Close()
	log.Println("LIFECYCLES LOADED: ", len(self.lifecycles))

}
