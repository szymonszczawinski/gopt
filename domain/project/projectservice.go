package project

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"
)

type IProjectService interface {
	GetProjects() []ProjectListItem
	GetProject(projectId string) (ProjectDetails, error)
	CreateProject(newProject CreateProjectCommand) (ProjectListItem, error)
}

type projectService struct {
	ctx        context.Context
	eg         *errgroup.Group
	repository IProjectRepository
}

func NewProjectService(eg *errgroup.Group, ctx context.Context, repository IProjectRepository) *projectService {
	instance := new(projectService)
	instance.repository = repository
	instance.ctx = ctx
	instance.eg = eg
	return instance
}

func (self *projectService) StartComponent() {

}

func (self projectService) GetProjects() []ProjectListItem {
	projects := self.repository.GetProjects()
	projectList := make([]ProjectListItem, 0)
	for _, project := range projects {
		projectList = append(projectList, NewProjectListItem(project))
	}
	return projectList
}

func (self projectService) GetProject(projectId string) (ProjectDetails, error) {
	project, err := self.repository.GetProject(projectId)
	if err != nil {
		return ProjectDetails{}, err
	}
	return NewProjectDetails(project), nil

}
func (self projectService) CreateProject(newProject CreateProjectCommand) (ProjectListItem, error) {
	projectLifecycle, err := self.repository.GetLifecycle()
	if err != nil {
		log.Println(err.Error())
		return ProjectListItem{}, err
	}
	project := NewProject(newProject.IssueKey, newProject.Name, projectLifecycle)
	stored, err := self.repository.StoreProject(project)
	if err != nil {
		log.Println("Could not create Project", err.Error())
		return ProjectListItem{}, err
	}
	return NewProjectListItem(stored), nil
}
