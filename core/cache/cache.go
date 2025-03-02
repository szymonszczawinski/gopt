package cache

import (
	"gopt/core/domain/common"
	"gopt/core/domain/project"
	"gopt/coreapi"
	"log/slog"
)

type IProjectQueryRepository interface {
	GetProjects() coreapi.Result[[]project.ProjectListElement]
}

type Cache struct {
	projectRepo       IProjectQueryRepository
	issueTypes        []common.IssueType
	availableProjects map[string]project.ProjectListElement
}

func NewCache(projectRepo IProjectQueryRepository) *Cache {
	return &Cache{
		projectRepo:       projectRepo,
		availableProjects: map[string]project.ProjectListElement{},
	}
}

func (c *Cache) InitCache() {
	slog.Info("cache load issue types")
	c.issueTypes = common.GetIssueTypes()
	slog.Info("cache load available projects")
	projectsResult := c.projectRepo.GetProjects()
	if projectsResult.Sucess() {
		for _, p := range projectsResult.Data() {
			c.availableProjects[p.ProjectKey] = p
		}
	}
}

func (c *Cache) AddProject(p project.ProjectListElement) {
	if _, ok := c.availableProjects[p.ProjectKey]; !ok {
		c.availableProjects[p.ProjectKey] = p
	}
}

func (c Cache) GetAvailabeProjects() []project.ProjectListElement {
	projects := []project.ProjectListElement{}
	for _, p := range c.availableProjects {
		projects = append(projects, p)
	}
	return projects
}

func (c Cache) GetIssueTypes() []common.IssueType {
	return c.issueTypes
}
