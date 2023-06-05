package storage

import (
	"gosi/issues"
)

type IStorage interface {
	GetProjects() []issues.Project
}
