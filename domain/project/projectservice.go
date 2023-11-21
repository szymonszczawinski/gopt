package project

import (
	"context"
	"gosi/coreapi"
	"log"

	"golang.org/x/sync/errgroup"
)

type IProjectService interface {
	GetProjects() coreapi.Result[[]ProjectListElement]
	GetProject(projectId string) coreapi.Result[ProjectDetails]
	CreateProject(newProject CreateProjectCommand) coreapi.Result[ProjectDetails]
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

func (service projectService) GetProjects() coreapi.Result[[]ProjectListElement] {
	return service.repository.GetProjects()
}

func (service projectService) GetProject(projectId string) coreapi.Result[ProjectDetails] {
	result := service.repository.GetProject(projectId)
	if !result.Sucess() {
		return coreapi.NewResult(ProjectDetails{}, result.Error())
	}
	return coreapi.NewResult(NewProjectDetails(result.Data()), nil)
}

func (service projectService) CreateProject(newProject CreateProjectCommand) coreapi.Result[ProjectDetails] {
	resultState := service.repository.GetProjectState()
	if !resultState.Sucess() {
		log.Println(resultState.Error())
		return coreapi.NewResult[ProjectDetails](ProjectDetails{}, resultState.Error())
	}
	project := NewProject(newProject.IssueKey, newProject.Name, resultState.Data())
	resultProject := service.repository.StoreProject(project)
	if !resultProject.Sucess() {
		log.Println("Could not create Project", resultProject.Error())
		return coreapi.NewResult[ProjectDetails](ProjectDetails{}, resultProject.Error())
	}
	return coreapi.NewResult[ProjectDetails](NewProjectDetails(resultProject.Data()), nil)
}

func (service *projectService) CloseProject(projectId string) coreapi.Result[ProjectDetails] {
	// TODO: to implement
	return coreapi.NewResult[ProjectDetails](ProjectDetails{}, nil)
}
