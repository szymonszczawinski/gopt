package bun

import (
	"context"
	"database/sql"
	"errors"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/sync/errgroup"
)

type bunDatabase struct {
	db  *bun.DB
	ctx context.Context
	eg  *errgroup.Group
}

func NewBunDatabase(eg *errgroup.Group, ctx context.Context) storage.IBunDatabase {
	instance := new(bunDatabase)
	instance.ctx = ctx
	instance.eg = eg
	return instance
}

func (self *bunDatabase) StartComponent() {
	log.Println("Starting", service.ComponentTypeBunDatabase)
	if databaseExists() {
		self.db, _ = openDatabase(storage.DatabaseDialectSqlite3)
	} else {
		self.db, _ = createAndInitDb(storage.DatabaseDialectSqlite3, self.ctx)
	}
}

func (self *bunDatabase) Close() {
	self.db.Close()
}

func (self *bunDatabase) NewSelect() *bun.SelectQuery {
	return self.db.NewSelect()
}

func (self *bunDatabase) NewInsert() *bun.InsertQuery {
	return self.db.NewInsert()

}

func databaseExists() bool {
	dbfile := os.Getenv("DATABASE_FILE_NAME")
	if _, err := os.Stat(dbfile); err != nil {
		log.Println("File", dbfile, " does not exists")
		return false
	}
	return true
}

func openDatabase(dialect storage.DatabaseDialect) (*bun.DB, error) {
	switch dialect {
	case storage.DatabaseDialectSqlite3:
		return mustOpenSqlite3Database(), nil
	case storage.DatabaseDialectMySql:
		return mustOpenMysqlDatabase(), nil
	case storage.DatabaseDialectPostgres:
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

func createAndInitDb(dialect storage.DatabaseDialect, ctx context.Context) (*bun.DB, error) {
	switch dialect {
	case storage.DatabaseDialectSqlite3:
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
