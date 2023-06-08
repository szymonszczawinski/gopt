package storage

import (
	"gosi/issues"
)

type IStorage interface {
	GetProjects() []issues.Project
	GetProject(projectId int64) (*issues.Project, error)
}
