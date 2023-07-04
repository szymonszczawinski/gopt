package bun

import (
	"context"
	"gosi/coreapi/storage"
	"gosi/issues/domain"
	"log"

	"golang.org/x/sync/errgroup"
)

type bunRepository struct {
}

func NewRepository(eg *errgroup.Group, ctx context.Context) *bunRepository {
	return new(bunRepository)
}
func (self *bunRepository) StartService() {
	log.Println("Starting", storage.IREPOSITORY)
}

func (self *bunRepository) GetProjects() []domain.Project {
	return nil
}
func (self *bunRepository) GetProject(projectId string) (domain.Project, error) {
	return domain.Project{}, nil
}
func (self *bunRepository) GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error) {
	return domain.Lifecycle{}, nil
}
func (self *bunRepository) StoreProject(project domain.Project) (domain.Project, error) {
	return domain.Project{}, nil
}
func (self *bunRepository) GetComments() []domain.Comment {
	return nil
}
func (self *bunRepository) StoreComment(comment domain.Comment) (domain.Comment, error) {
	return domain.Comment{}, nil
}
