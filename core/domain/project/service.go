package project

import (
	"context"
	"errors"
	"gopt/coreapi"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

var ErrorCouldNotInsertProject = errors.New("could not insert project")

type IProjectRepository interface {
	coreapi.IComponent
	GetProjects() coreapi.Result[[]ProjectListElement]
	GetProject(projectId string) coreapi.Result[Project]
	StoreProject(project Project) coreapi.Result[Project]
	UpdateProject(project Project) coreapi.Result[Project]
}
type (
	IProjectCache interface {
		AddProject(p ProjectListElement)
	}
	projectService struct {
		ctx        context.Context
		eg         *errgroup.Group
		repository IProjectRepository
		cache      IProjectCache
	}
)

func NewProjectService(eg *errgroup.Group, ctx context.Context, repository IProjectRepository, cache IProjectCache) projectService {
	instance := projectService{
		repository: repository,
		cache:      cache,
		ctx:        ctx,
		eg:         eg,
	}
	return instance
}

func (service projectService) StartComponent() {
}

// func (service projectService) GetProjects() coreapi.Result[[]ProjectListElement] {
// 	return service.repository.GetProjects()
// }

func (service projectService) GetProject(command GetProject) coreapi.Result[ProjectDetails] {
	result := service.repository.GetProject(command.ProjectKey)
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

func (service projectService) CreateProject(command CreateProject) coreapi.Result[ProjectDetails] {
	newProjectState := NewProjectState(1, 1, "Open")
	project := NewProject(command.ProjectKey, command.Name, newProjectState)

	resultProject := service.repository.StoreProject(project)
	if !resultProject.Sucess() {
		slog.Info("Could not create Project", "err", resultProject.Error())
		return coreapi.NewResult(ProjectDetails{}, resultProject.Error())
	}
	p := resultProject.Data()
	service.cache.AddProject(ProjectListElement{
		ProjectKey: p.projectKey,
		Name:       p.name,
		Id:         p.Id,
	})
	return coreapi.NewResult(NewProjectDetails(p), nil)
}

func (service projectService) CloseProject(projectId string) coreapi.Result[ProjectDetails] {
	// TODO: to implement
	return coreapi.NewResult(ProjectDetails{}, nil)
}
