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
	CloseProject(projectId string) (ProjectDetails, error)
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

func (service *projectService) StartComponent() {

}

func (service projectService) GetProjects() []ProjectListItem {
	projects := service.repository.GetProjects()
	projectList := make([]ProjectListItem, 0)
	for _, project := range projects {
		projectList = append(projectList, NewProjectListItem(project))
	}
	return projectList
}

func (service projectService) GetProject(projectId string) (ProjectDetails, error) {
	project, err := service.repository.GetProject(projectId)
	if err != nil {
		return ProjectDetails{}, err
	}
	return NewProjectDetails(project), nil

}
func (service projectService) CreateProject(newProject CreateProjectCommand) (ProjectListItem, error) {
	projectState, err := service.repository.GetProjectState()
	if err != nil {
		log.Println(err.Error())
		return ProjectListItem{}, err
	}
	project := NewProject(newProject.IssueKey, newProject.Name, projectState)
	stored, err := service.repository.StoreProject(project)
	if err != nil {
		log.Println("Could not create Project", err.Error())
		return ProjectListItem{}, err
	}
	return NewProjectListItem(stored), nil
}

func (service *projectService) CloseProject(projectId string) (ProjectDetails, error) {
	//TODO: to implement
	return ProjectDetails{}, nil
}
