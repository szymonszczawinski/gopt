package memory

import (
	"errors"
	"fmt"
	"gosi/coreapi/storage"
	"gosi/issues/domain"
	"log"

	"golang.org/x/exp/maps"
)

type idsRepository struct {
	projectId int
	bugId     int
	commentId int
}

func (self *idsRepository) getNextProjectId() int {
	self.projectId += 1
	return self.projectId
}
func (self *idsRepository) getNextCommentId() int {
	log.Println("Current CommentId:", self.commentId)

	self.commentId += 1
	log.Println("New CommentId:", self.commentId)
	return self.commentId
}

type memoryRepository struct {
	idsRepository
	lifecycleStates map[int]domain.LifecycleState
	lifecycles      map[int]domain.Lifecycle
	projects        map[string]domain.Project
	comments        map[int]domain.Comment
}

func (s *memoryRepository) initStorage() {
	s.initLifecycleStates()
	s.initLifecycles()
	s.initProjects()
	s.comments = map[int]domain.Comment{}
}

func (s *memoryRepository) initLifecycleStates() {
	s.lifecycleStates = make(map[int]domain.LifecycleState)
	s.lifecycleStates[1] = domain.NewLifecycleState(1, domain.LIFECYCLE_STATE_DRAFT)
	s.lifecycleStates[2] = domain.NewLifecycleState(2, domain.LIFECYCLE_STATE_NEW)
	s.lifecycleStates[3] = domain.NewLifecycleState(3, domain.LIFECYCLE_STATE_ANALISYS)
	s.lifecycleStates[4] = domain.NewLifecycleState(4, domain.LIFECYCLE_STATE_DESIGN)
	s.lifecycleStates[5] = domain.NewLifecycleState(5, domain.LIFECYCLE_STATE_DEVELOPMENT)
	s.lifecycleStates[6] = domain.NewLifecycleState(6, domain.LIFECYCLE_STATE_OPEN)
	s.lifecycleStates[7] = domain.NewLifecycleState(7, domain.LIFECYCLE_STATE_CLOSED)
	s.lifecycleStates[8] = domain.NewLifecycleState(8, domain.LIFECYCLE_STATE_INTEGRATION)
	s.lifecycleStates[9] = domain.NewLifecycleState(9, domain.LIFECYCLE_STATE_VERIFICATION)
	s.lifecycleStates[10] = domain.NewLifecycleState(10, domain.LIFECYCLE_STATE_FIXED)
	s.lifecycleStates[11] = domain.NewLifecycleState(11, domain.LIFECYCLE_STATE_CLOSED)
	s.lifecycleStates[12] = domain.NewLifecycleState(12, domain.LIFECYCLE_STATE_REJECTED)
	s.lifecycleStates[13] = domain.NewLifecycleState(13, domain.LIFECYCLE_STATE_RETEST)
}

func (s *memoryRepository) initLifecycles() {
	s.lifecycles = make(map[int]domain.Lifecycle)
	//DRAFT -> NEW -> ANALISYS -> DESIGN -> DEV -> CLOSED
	s.lifecycles[1] = domain.NewLifeCycleBuilder(1, "Project", s.lifecycleStates[1]).
		AddTransition(s.lifecycleStates[1], s.lifecycleStates[2]).
		AddTransition(s.lifecycleStates[2], s.lifecycleStates[3]).
		AddTransition(s.lifecycleStates[3], s.lifecycleStates[4]).
		AddTransition(s.lifecycleStates[4], s.lifecycleStates[5]).
		AddTransition(s.lifecycleStates[7], s.lifecycleStates[7]).
		Build()
}
func (s *memoryRepository) initProjects() {
	s.projects = make(map[string]domain.Project)
	projectA := domain.NewProject("PROJA", "Project A", s.lifecycles[1])
	projectA.SetId(1)
	s.projects[projectA.GetItemKey()] = projectA

}

func NewMemoryRepository() storage.IRepository {
	instance := memoryRepository{
		idsRepository: idsRepository{},
	}

	return &instance

}

func (self *memoryRepository) StartService() {
	log.Println("Starting", storage.IREPOSITORY)
	self.initStorage()
}

func (s memoryRepository) GetProjects() []domain.Project {
	return maps.Values(s.projects)
}

func (self memoryRepository) GetProject(projectId string) (domain.Project, error) {
	project, exist := self.projects[projectId]
	if exist {
		return project, nil
	}
	return domain.Project{}, errors.New(fmt.Sprintf("Project with ID: %v not found", projectId))
}
func (self memoryRepository) GetLifecycle(issueType domain.IssueType) (domain.Lifecycle, error) {
	if issueType == domain.TProject {
		return self.lifecycles[1], nil
	}
	return domain.Lifecycle{}, errors.New(fmt.Sprintf("No Lifecycle defined for issue type: %v", issueType))
}
func (self memoryRepository) StoreProject(project domain.Project) (domain.Project, error) {
	if _, exists := self.projects[project.GetItemKey()]; exists {
		return domain.Project{}, errors.New(fmt.Sprintf("Project already exists"))
	}
	project.SetId(self.getNextCommentId())
	self.projects[project.GetItemKey()] = project
	return project, nil
}

func (self memoryRepository) GetComments() []domain.Comment {
	return maps.Values(self.comments)
}

func (self *memoryRepository) StoreComment(comment domain.Comment) (domain.Comment, error) {
	comment.SetId(self.getNextCommentId())
	self.comments[comment.GetId()] = comment
	//TODO: Move adding comment and updating project to Project service
	for _, project := range self.projects {
		if project.GetId() == comment.GetParentItemId() {
			project.AddComment(comment)
			self.projects[project.GetItemKey()] = project
			break
		}

	}
	return comment, nil
}
