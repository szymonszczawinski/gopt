package bun

import (
	"context"
	"database/sql"
	"errors"
	"gosi/core/storage/dao"
	"gosi/issues/domain"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DatabaseDialect string

const (
	DatabaseDialectSqlite3  DatabaseDialect = "sqlite3"
	DatabaseDialectMySql    DatabaseDialect = "mysql"
	DatabaseDialectPostgres DatabaseDialect = "postgres"
)

func databaseExists() bool {
	dbfile := os.Getenv("DATABASE_FILE_NAME")
	if _, err := os.Stat(dbfile); err != nil {
		log.Println("File", dbfile, " does not exists")
		return false
	}
	return true
}

func openDatabase(dialect DatabaseDialect) (*bun.DB, error) {
	switch dialect {
	case DatabaseDialectSqlite3:
		return mustOpenSqlite3Database(), nil
	case DatabaseDialectMySql:
		return mustOpenMysqlDatabase(), nil
	case DatabaseDialectPostgres:
		return mustOpenPostgresDatabase(), nil
	}
	return nil, errors.New("Could not open database")
}

func mustOpenPostgresDatabase() *bun.DB {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	return bun.NewDB(sqldb, pgdialect.New())
}

func mustOpenMysqlDatabase() *bun.DB {
	connectionString := os.Getenv("CONNECTION_STRING")
	sqldb, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}

	return bun.NewDB(sqldb, mysqldialect.New())
}

func mustOpenSqlite3Database() *bun.DB {
	dbfile := os.Getenv("DATABASE_FILE_NAME")
	sqldb, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Panic(err.Error())
	}
	return bun.NewDB(sqldb, sqlitedialect.New())
}

func createAndInitDb(dialect DatabaseDialect, ctx context.Context) (*bun.DB, error) {
	switch dialect {
	case DatabaseDialectSqlite3:
		return mustCreateSqlite3Database(ctx), nil
	}
	return nil, nil
}

func mustCreateSqlite3Database(ctx context.Context) *bun.DB {
	dbfile := os.Getenv("DATABASE_FILE_NAME")
	sqldb, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Panic(err.Error())
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	mustInitDatabase(db, ctx)
	return db
}

func mustInitDatabase(db *bun.DB, ctx context.Context) {
	_, err := db.NewCreateTable().
		Model((*dao.LifecycleStateRow)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		panic(err)
	}
	states := []dao.LifecycleStateRow{
		{Id: 1, Name: "Draft"},
		{Id: 2, Name: "New"},
		{Id: 3, Name: "Open"},
		{Id: 4, Name: "Analysis"},
		{Id: 5, Name: "Design"},
		{Id: 6, Name: "Development"},
		{Id: 7, Name: "Integration"},
		{Id: 8, Name: "Verification"},
		{Id: 9, Name: "Fixed"},
		{Id: 10, Name: "Closed"},
		{Id: 11, Name: "Retest"},
		{Id: 12, Name: "Rejected"},
		{Id: 13, Name: "Assigned"},
	}

	_, err = db.NewInsert().Model(&states).Exec(ctx)
	if err != nil {
		panic(err)
	}
	_, err = db.NewCreateTable().
		Model((*dao.LifecycleRow)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	lifecycle := dao.LifecycleRow{Id: 1, Name: "Project", StartStateId: 1}
	_, err = db.NewInsert().Model(&lifecycle).Exec(ctx)
	if err != nil {
		panic(err)
	}
	_, err = db.NewCreateTable().
		Model((*dao.StateTransition)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		panic(err)
	}
	//Project :: DRAFT -> NEW -> ANALISYS -> DESIGN -> DEV -> CLOSED
	transitions := []dao.StateTransition{
		{LifecycleId: 1, FromStateId: 1, ToStateId: 2},
		{LifecycleId: 1, FromStateId: 2, ToStateId: 4},
		{LifecycleId: 1, FromStateId: 4, ToStateId: 5},
		{LifecycleId: 1, FromStateId: 5, ToStateId: 6},
		{LifecycleId: 1, FromStateId: 6, ToStateId: 10},
	}
	_, err = db.NewInsert().Model(&transitions).Exec(ctx)
	if err != nil {
		panic(err)
	}
	_, err = db.NewCreateTable().
		Model((*dao.ProjectRow)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		panic(err)
	}
	project := dao.ProjectRow{
		Name:        "COSMOS",
		ItemKey:     "COSMOS",
		ItemNumber:  1,
		Description: "Super COSMOS Project",
		StateId:     1,
		LifecycleId: 1,
		CreatedById: 1,
	}
	_, err = db.NewInsert().Model(&project).Exec(ctx)
	if err != nil {
		panic(err)
	}

}
func (self *bunRepository) loadDictionaryData() {
	self.loadLifecycles()
}

func (self *bunRepository) loadLifecycles() {
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
