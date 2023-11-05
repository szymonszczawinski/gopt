package project

import (
	"context"
	"gosi/coreapi"
	"log"

	"golang.org/x/sync/errgroup"
)

type IProjectService interface {
	GetProjects() []ProjectListItem
	GetProject(projectId string) coreapi.Result[ProjectDetails]
	CreateProject(newProject CreateProjectCommand) coreapi.Result[ProjectListItem]
	CloseProject(projectId string) coreapi.Result[ProjectDetails]
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

func (service projectService) GetProject(projectId string) coreapi.Result[ProjectDetails] {
	result := service.repository.GetProject(projectId)
	if !result.Sucess() {
		return coreapi.NewResult(ProjectDetails{}, result.Error())

	}
	return coreapi.NewResult(NewProjectDetails(result.Data()), nil)

}
func (service projectService) CreateProject(newProject CreateProjectCommand) coreapi.Result[ProjectListItem] {
	resultState := service.repository.GetProjectState()
	if !resultState.Sucess() {
		log.Println(resultState.Error())
		return coreapi.NewResult[ProjectListItem](ProjectListItem{}, resultState.Error())
	}
	project := NewProject(newProject.IssueKey, newProject.Name, resultState.Data())
	resultProject := service.repository.StoreProject(project)
	if !resultProject.Sucess() {
		log.Println("Could not create Project", resultProject.Error())
		return coreapi.NewResult[ProjectListItem](ProjectListItem{}, resultProject.Error())
	}
	return coreapi.NewResult[ProjectListItem](NewProjectListItem(resultProject.Data()), nil)
}

func (service *projectService) CloseProject(projectId string) coreapi.Result[ProjectDetails] {
	//TODO: to implement
	return coreapi.NewResult[ProjectDetails](ProjectDetails{}, nil)
}
