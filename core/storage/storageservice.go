package storage

import (
	"context"
	"gosi/coreapi/queue"
	"gosi/issues"
	"log"

	"golang.org/x/sync/errgroup"
)

const (
	ISTORAGESERVICE = "IStorageService"
)

type IStorageService interface {
	GetProjects() []issues.Project
	GetProject(projectId int64) (*issues.Project, error)
}

type storageService struct {
	ctx     context.Context
	looper  queue.IJobQueue
	storage IStorage
}

func NewStorageService(eg *errgroup.Group, ctx context.Context) *storageService {
	log.Println("New Storage Service")
	instance := new(storageService)
	instance.ctx = ctx
	instance.looper = queue.NeqJobQueue("storageService", eg)
	instance.storage = GetStorage()
	return instance
}
func (s *storageService) StartService() {
	log.Println("Starting", ISTORAGESERVICE)
	s.looper.Start(s.ctx)
}

func (s storageService) GetProjects() []issues.Project {
	resultChan := make(chan []issues.Project)
	s.looper.Add(&queue.Job{Execute: func() {
		log.Println("New Job :: getProjetcs")
		resultChan <- s.storage.GetProjects()
	}})

	return <-resultChan
}
func (s storageService) GetProject(projectId int64) (*issues.Project, error) {
	resultChan := make(chan issues.Project)
	errChan := make(chan error)
	s.looper.Add(&queue.Job{Execute: func() {
		project, err := s.storage.GetProject(projectId)
		if err != nil {
			errChan <- err
		} else {
			resultChan <- *project
		}
	}})
	select {
	case err := <-errChan:
		return nil, err
	case res := <-resultChan:
		return &res, nil
	}
}
