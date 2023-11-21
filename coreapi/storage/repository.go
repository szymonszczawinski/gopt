package storage

import (
	"gosi/coreapi/service"

	"github.com/uptrace/bun"
)

type (
	RepositoryType  int
	DatabaseDialect string
)

const (
	RepositoryTypeMemory    RepositoryType  = 1
	RepositoryTypeSql       RepositoryType  = 2
	RepositoryTypeBun       RepositoryType  = 3
	DatabaseDialectSqlite3  DatabaseDialect = "sqlite3"
	DatabaseDialectMySql    DatabaseDialect = "mysql"
	DatabaseDialectPostgres DatabaseDialect = "postgres"
)

type IBunDatabase interface {
	service.IComponent
	NewSelect() *bun.SelectQuery
	NewInsert() *bun.InsertQuery
	NewRaw(sql string, args ...interface{}) *bun.RawQuery
}
