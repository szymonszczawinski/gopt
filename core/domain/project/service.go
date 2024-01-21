package project

import (
	"context"
	"gopt/coreapi"
	"gopt/coreapi/service"
	"log"

	"golang.org/x/sync/errgroup"
)

type IProjectService interface {
	service.IComponent
	// GetProjects() coreapi.Result[[]ProjectListElement]
	GetProject(projectId string) coreapi.Result[ProjectDetails]
	CreateProject(newProject CreateProjectCommand) coreapi.Result[ProjectDetails]
	CloseProject(projectId string) coreapi.Result[ProjectDetails]
}

type projectService struct {
	ctx        context.Context
	eg         *errgroup.Group
	repository IProjectRepository
}

func NewProjectService(eg *errgroup.Group, ctx context.Context, repository IProjectRepository) IProjectService {
	instance := new(projectService)
	instance.repository = repository
	instance.ctx = ctx
	instance.eg = eg
	return instance
}

func (service *projectService) StartComponent() {
}

// func (service projectService) GetProjects() coreapi.Result[[]ProjectListElement] {
// 	return service.repository.GetProjects()
// }

func (service projectService) GetProject(projectId string) coreapi.Result[ProjectDetails] {
	result := service.repository.GetProject(projectId)
	// FIXME: remove log
	log.Println("projectService GetProject", result)
	if !result.Sucess() {
		return coreapi.NewResult(ProjectDetails{}, result.Error())
	}
	return coreapi.NewResult(NewProjectDetails(result.Data()), nil)
}

func (service projectService) CreateProject(newProject CreateProjectCommand) coreapi.Result[ProjectDetails] {
	newProjectState := NewProjectState(1, 1, "Open")
	project := NewProject(newProject.IssueKey, newProject.Name, newProjectState)

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
