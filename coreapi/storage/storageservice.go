package storage

import "gosi/issues/domain"

const (
	ISTORAGESERVICE = "IStorageService"
)

type IStorageService interface {
	CreateProject(project domain.Project) (domain.Project, error)
	CreateComment(comment domain.Comment) (domain.Comment, error)
}
