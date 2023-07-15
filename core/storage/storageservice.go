package storage

import (
	"context"
	"gosi/coreapi/queue"
	"gosi/coreapi/service"
	"gosi/coreapi/storage"
	"gosi/issues/domain"
	"log"

	"golang.org/x/sync/errgroup"
)

type storageService struct {
	ctx        context.Context
	looper     queue.IJobQueue
	repository storage.IRepository
}

func NewStorageService(eg *errgroup.Group, ctx context.Context, repository storage.IRepository) storage.IStorageService {
	log.Println("New Storage Service")
	instance := new(storageService)
	instance.ctx = ctx
	instance.looper = queue.NeqJobQueue("storageService", eg)
	instance.repository = repository
	return instance
}
func (self *storageService) StartComponent() {
	log.Println("Starting", service.ComponentTypeTypeStorageService)
	self.looper.Start(self.ctx)
}

func (self storageService) CreateProject(project domain.Project) (domain.Project, error) {
	errChan := make(chan error)
	resChan := make(chan domain.Project)

	defer close(errChan)
	defer close(resChan)
	self.looper.Add(&queue.Job{Execute: func() {
		stored, err := self.repository.StoreProject(project)
		if err != nil {
			errChan <- err
		} else {
			resChan <- stored
		}

	}})
	select {
	case err := <-errChan:
		return domain.Project{}, err
	case stored := <-resChan:
		return stored, nil
	}

}
func (self storageService) CreateComment(comment domain.Comment) (domain.Comment, error) {
	errChan := make(chan error)
	resChan := make(chan domain.Comment)
	defer close(errChan)
	defer close(resChan)
	self.looper.Add(&queue.Job{Execute: func() {
		stored, err := self.repository.StoreComment(comment)
		if err != nil {
			errChan <- err
		} else {
			resChan <- stored
		}

	}})
	select {
	case err := <-errChan:
		return domain.Comment{}, err
	case stored := <-resChan:
		return stored, nil
	}

}
