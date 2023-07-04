package storage

import "gosi/issues/domain"

type RepositoryType int

const (
	IREPOSITORY                = "IRepository"
	MEMORY      RepositoryType = 1
	SQL         RepositoryType = 2
	BUN         RepositoryType = 3
)

type IRepository interface {
	GetProjects() []domain.Project
	GetProject(projectId string) (domain.Project, error)
	GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error)
	StoreProject(project domain.Project) (domain.Project, error)
	GetComments() []domain.Comment
	StoreComment(comment domain.Comment) (domain.Comment, error)
}
