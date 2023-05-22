package storage

import (
	"gosi/model"
)

var instance *storage

type IStorage interface {
	GetProjects() []model.Project
}

type storage struct {
	projects []model.Project
}

func (s *storage) initStorage() {
	s.projects = append(s.projects, model.Project{
		Name:   "Project A",
		Id:     1234,
		Status: model.ProjectStatus_NEW,
		Name2:  "Super Project",
	})
	s.projects = append(s.projects, model.Project{
		Name:   "Project B",
		Id:     1235,
		Status: model.ProjectStatus_NEW,
		Name2:  "Super Project 2",
	})
}

func (s *storage) GetProjects() []model.Project {
	return s.projects
}

func GetStorage() IStorage {
	if instance == nil {
		instance = new(storage)
		instance.projects = make([]model.Project, 0)

	}
	return instance
}
