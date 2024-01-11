package storage

type (
	RepositoryType  int
	DatabaseDialect string
)

const (
	RepositoryTypeMemory    RepositoryType  = 1
	RepositoryTypeSql       RepositoryType  = 2
	DatabaseDialectSqlite3  DatabaseDialect = "sqlite3"
	DatabaseDialectMySql    DatabaseDialect = "mysql"
	DatabaseDialectPostgres DatabaseDialect = "postgres"
)
