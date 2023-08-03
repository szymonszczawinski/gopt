package service

import (
	"context"
	"gosi/issues/domain"
	"gosi/issues/dto"
	"gosi/issues/storage"
	"log"

	"golang.org/x/sync/errgroup"
)

type IProjectService interface {
	GetProjects() []dto.ProjectListItem
	GetProject(projectId string) (dto.ProjectDetails, error)
	CreateProject(newProject dto.CreateProjectCommand) (dto.ProjectListItem, error)
	AddComment(newComment dto.AddCommentCommand) (domain.Comment, error)
}

type projectService struct {
	ctx        context.Context
	eg         *errgroup.Group
	repository storage.IIssueRepository
}

func NewProjectService(eg *errgroup.Group, ctx context.Context, repository storage.IIssueRepository) *projectService {
	instance := new(projectService)
	instance.repository = repository
	instance.ctx = ctx
	instance.eg = eg
	return instance
}

func (self *projectService) StartComponent() {

}

func (self projectService) GetProjects() []dto.ProjectListItem {
	projects := self.repository.GetProjects()
	projectList := make([]dto.ProjectListItem, 0)
	for _, project := range projects {
		projectList = append(projectList, dto.NewProjectListItem(project))
	}
	return projectList
}

func (self projectService) GetProject(projectId string) (dto.ProjectDetails, error) {
	project, err := self.repository.GetProject(projectId)
	if err != nil {
		return dto.ProjectDetails{}, err
	}
	return dto.NewProjectDetails(project), nil

}
func (self projectService) CreateProject(newProject dto.CreateProjectCommand) (dto.ProjectListItem, error) {
	projectLifecycle, err := self.repository.GetLifecycle(domain.IssueTypeProject)
	if err != nil {
		log.Println(err.Error())
		return dto.ProjectListItem{}, err
	}
	project := domain.NewProject(newProject.IssueKey, newProject.Name, projectLifecycle)
	stored, err := self.repository.StoreProject(project)
	if err != nil {
		log.Println("Could not create Project", err.Error())
		return dto.ProjectListItem{}, err
	}
	return dto.NewProjectListItem(stored), nil
}

func (self projectService) AddComment(newComment dto.AddCommentCommand) (domain.Comment, error) {
	project, err := self.repository.GetProject(newComment.ParentIssueKey)
	if err != nil {
		return domain.Comment{}, err
	}
	stored, err := self.repository.StoreComment(domain.NewComment(project.GetId(), newComment.Content))
	if err != nil {
		return domain.Comment{}, err
	} else {
		return stored, nil
	}

}
