package storage

import "gosi/issues/domain"

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

type IRepository interface {
	GetProjects() []domain.Project
	GetProject(projectId string) (domain.Project, error)
	GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error)
	StoreProject(project domain.Project) (domain.Project, error)
	GetComments() []domain.Comment
	StoreComment(comment domain.Comment) (domain.Comment, error)
}
