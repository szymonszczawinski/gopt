package project

import (
	"context"
	"gopt/coreapi"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

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

// func (service projectService) GetProjects() coreapi.Result[[]ProjectListElement] {
// 	return service.repository.GetProjects()
// }

func (service projectService) GetProject(projectId string) coreapi.Result[ProjectDetails] {
	result := service.repository.GetProject(projectId)
	// FIXME: remove log
	slog.Info("projectService GetProject", "result", result)
	if !result.Sucess() {
		return coreapi.NewResult(ProjectDetails{}, result.Error())
	}

	projectDetails := NewProjectDetails(result.Data())
	projectDetails.Items = coreapi.Map2(result.Data().items,
		func(t ProjectItem) ProjectDetailsItem {
			return ProjectDetailsItem{
				State:      t.itemState,
				ItemType:   string(t.itemType),
				Name:       t.name,
				ItemKey:    t.itemKey,
				Created:    t.GetCreationTime().String(),
				Updated:    t.GetLastUpdateTime().String(),
				AssignedTo: t.assignee,
			}
		},
	)
	return coreapi.NewResult(projectDetails, nil)
}

func (service projectService) CreateProject(newProject CreateProjectCommand) coreapi.Result[ProjectDetails] {
	newProjectState := NewProjectState(1, 1, "Open")
	project := NewProject(newProject.IssueKey, newProject.Name, newProjectState)

	resultProject := service.repository.StoreProject(project)
	if !resultProject.Sucess() {
		slog.Info("Could not create Project", "err", resultProject.Error())
		return coreapi.NewResult(ProjectDetails{}, resultProject.Error())
	}
	return coreapi.NewResult(NewProjectDetails(resultProject.Data()), nil)
}

func (service *projectService) CloseProject(projectId string) coreapi.Result[ProjectDetails] {
	// TODO: to implement
	return coreapi.NewResult(ProjectDetails{}, nil)
}
