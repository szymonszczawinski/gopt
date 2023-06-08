package storage

import "gosi/issues"

var instance *memoryStorage

type memoryStorage struct {
	lifecycleStates map[int]issues.LifecycleState
	lifecycles      map[int]issues.Lifecycle
	projects        []issues.Project
}

func (s *memoryStorage) initStorage() {
	s.initLifecycleStates()
	s.initLifecycles()
	s.initProjects()
}

func (s *memoryStorage) initLifecycleStates() {
	s.lifecycleStates = make(map[int]issues.LifecycleState)
	s.lifecycleStates[1] = issues.NewLifecycleState(1, issues.LIFECYCLE_STATE_DRAFT)
	s.lifecycleStates[2] = issues.NewLifecycleState(2, issues.LIFECYCLE_STATE_NEW)
	s.lifecycleStates[3] = issues.NewLifecycleState(3, issues.LIFECYCLE_STATE_ANALISYS)
	s.lifecycleStates[4] = issues.NewLifecycleState(4, issues.LIFECYCLE_STATE_DESIGN)
	s.lifecycleStates[5] = issues.NewLifecycleState(5, issues.LIFECYCLE_STATE_DEVELOPMENT)
	s.lifecycleStates[6] = issues.NewLifecycleState(6, issues.LIFECYCLE_STATE_OPEN)
	s.lifecycleStates[7] = issues.NewLifecycleState(7, issues.LIFECYCLE_STATE_CLOSED)
	s.lifecycleStates[8] = issues.NewLifecycleState(8, issues.LIFECYCLE_STATE_INTEGRATION)
	s.lifecycleStates[9] = issues.NewLifecycleState(9, issues.LIFECYCLE_STATE_VERIFICATION)
	s.lifecycleStates[10] = issues.NewLifecycleState(10, issues.LIFECYCLE_STATE_FIXED)
	s.lifecycleStates[11] = issues.NewLifecycleState(11, issues.LIFECYCLE_STATE_CLOSED)
	s.lifecycleStates[12] = issues.NewLifecycleState(12, issues.LIFECYCLE_STATE_REJECTED)
	s.lifecycleStates[13] = issues.NewLifecycleState(13, issues.LIFECYCLE_STATE_RETEST)
}

func (s *memoryStorage) initLifecycles() {
	s.lifecycles = make(map[int]issues.Lifecycle)
	s.lifecycles[1] = issues.NewLifeCycleBuilder(1, "Project", s.lifecycleStates[1]).
		AddTransition(s.lifecycleStates[1], s.lifecycleStates[3]).AddTransition(s.lifecycleStates[3], s.lifecycleStates[4]).
		AddTransition(s.lifecycleStates[4], s.lifecycleStates[5]).AddTransition(s.lifecycleStates[5], s.lifecycleStates[8]).
		AddTransition(s.lifecycleStates[8], s.lifecycleStates[9]).AddTransition(s.lifecycleStates[9], s.lifecycleStates[11]).
		Build()
}
func (s *memoryStorage) initProjects() {
	s.projects = make([]issues.Project, 0)
	s.projects = append(s.projects, *issues.NewProject(1, "Project A", s.lifecycles[1]))

}

func (s *memoryStorage) GetProjects() []issues.Project {
	return s.projects
}

func GetStorage() IStorage {
	if instance == nil {
		instance = new(memoryStorage)
		instance.initStorage()
	}
	return instance
}
