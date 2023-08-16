package storage

import (
	"github.com/uptrace/bun"
	"gosi/coreapi/service"
)

type RepositoryType int
type DatabaseDialect string

const (
	RepositoryTypeMemory RepositoryType = 1
	RepositoryTypeSql    RepositoryType = 2
	RepositoryTypeBun    RepositoryType = 3
)

const (
	DatabaseDialectSqlite3  DatabaseDialect = "sqlite3"
	DatabaseDialectMySql    DatabaseDialect = "mysql"
	DatabaseDialectPostgres DatabaseDialect = "postgres"
)

type IBunDatabase interface {
	service.IComponent
	NewSelect() *bun.SelectQuery
	NewInsert() *bun.InsertQuery
}
