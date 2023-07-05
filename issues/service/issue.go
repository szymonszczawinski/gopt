package service

import (
	"errors"
	"gosi/core/logger"
	"gosi/core/service"
	iservice "gosi/coreapi/service"
	"gosi/coreapi/storage"
	"gosi/issues/domain"
	"gosi/issues/dto"
	"log"
)

type ProjectService struct {
	storageService storage.IStorageService
	repository     storage.IRepository
	logger         logger.Log
}

func NewProjectService() *ProjectService {
	instance := new(ProjectService)
	instance.storageService, _ = getStorageService()
	instance.repository, _ = getRepository()
	instance.logger = logger.NewLogger()
	return instance
}

func (self ProjectService) GetProjects() []dto.ProjectListItem {
	projects := self.repository.GetProjects()
	projectList := make([]dto.ProjectListItem, 0)
	for _, project := range projects {
		projectList = append(projectList, dto.NewProjectListItem(project))
	}
	return projectList
}

func (self ProjectService) GetProject(projectId string) (dto.ProjectDetails, error) {
	project, err := self.repository.GetProject(projectId)
	if err != nil {
		return dto.ProjectDetails{}, err
	}
	return dto.NewProjectDetails(project), nil

}
func (self ProjectService) CreateProject(newProject dto.CreateProjectCommand) (dto.ProjectDetails, error) {
	projectLifecycle, err := self.repository.GetLifecycle(domain.TProject)
	if err != nil {
		self.logger.Error(err.Error())
		return dto.ProjectDetails{}, err
	}
	project := domain.NewProject(newProject.IssueKey, newProject.Name, projectLifecycle)
	stored, err := self.storageService.CreateProject(project)
	if err != nil {
		self.logger.Error("Could not create Project", err.Error())
		return dto.ProjectDetails{}, err
	}
	return dto.NewProjectDetails(stored), nil
}

func (self ProjectService) AddComment(newComment dto.AddCommentCommand) (domain.Comment, error) {
	project, err := self.repository.GetProject(newComment.ParentIssueKey)
	if err != nil {
		return domain.Comment{}, err
	}
	stored, err := self.storageService.CreateComment(domain.NewComment(project.GetId(), newComment.Content))
	if err != nil {
		return domain.Comment{}, err
	} else {
		return stored, nil
	}

}

func getStorageService() (storage.IStorageService, error) {

	serviceManager, err := service.GetServiceManager()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	service, err := serviceManager.GetService(iservice.ServiceTypeIStorageService)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	storageService, ok := service.(storage.IStorageService)
	if !ok {
		return nil, errors.New("StorageService has incorrect type")
	}
	return storageService, nil
}
func getRepository() (storage.IRepository, error) {

	serviceManager, err := service.GetServiceManager()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	service, err := serviceManager.GetService(iservice.ServiceTypeIRepository)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	repository, ok := service.(storage.IRepository)
	if !ok {
		return nil, errors.New("Repository has incorrect type")
	}
	return repository, nil
}
