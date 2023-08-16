package storage

import "gosi/project/domain"

type IStorageService interface {
	CreateProject(project domain.Project) (domain.Project, error)
	CreateComment(comment domain.Comment) (domain.Comment, error)
}
