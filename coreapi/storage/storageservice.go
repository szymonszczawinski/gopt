package storage

import "gosi/issues/domain"

type IStorageService interface {
	CreateProject(project domain.Project) (domain.Project, error)
	CreateComment(comment domain.Comment) (domain.Comment, error)
}
