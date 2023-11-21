package bun

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
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
	dialect := getDialect()
	log.Println("Dialect:", dialect)
	self.db, _ = openDatabase(dialect, self.ctx)
}

func getDialect() storage.DatabaseDialect {
	dialectString := os.Getenv("DB_DIALECT")
	switch dialectString {
	case "postgres":
		return storage.DatabaseDialectPostgres
	case "slite3":
		return storage.DatabaseDialectSqlite3
	}
	panic(fmt.Sprintf("Dialect %v not supported", dialectString))
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

func (db *bunDatabase) NewRaw(sql string, args ...interface{}) *bun.RawQuery {
	return db.db.NewRaw(sql, args)
}

func openDatabase(dialect storage.DatabaseDialect, ctx context.Context) (*bun.DB, error) {
	switch dialect {
	case storage.DatabaseDialectPostgres:
		return mustOpenAndInitPostgresDb(ctx), nil
	}
	return nil, errors.New("Could not open database")
}

func mustOpenPostgresDatabase() *bun.DB {
	dsn := os.Getenv("DB_URL")
	log.Println("DSN", dsn)
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(os.Getenv("DB_HOST_PORT")),
		pgdriver.WithUser(os.Getenv("DB_USER")),
		pgdriver.WithPassword(os.Getenv("DB_PASS")),
		pgdriver.WithDatabase(os.Getenv("DB_NAME")),
	)
	sqldb := sql.OpenDB(pgconn)
	bundb := bun.NewDB(sqldb, pgdialect.New())
	// bundb.AddQueryHook(bundebug.NewQueryHook(
	// bundebug.WithVerbose(true),
	// bundebug.FromEnv("BUNDEBUG"),
	// ))
	return bundb
}

func mustOpenAndInitPostgresDb(ctx context.Context) *bun.DB {
	log.Println("Open DB")
	db := mustOpenPostgresDatabase()
	log.Println("Init DB")
	mustInitDatabase(db, ctx)
	return db
}
