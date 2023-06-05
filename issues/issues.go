package issues

type Issue struct {
	id           int64
	name         string
	description  string
	currentState LifecycleState
	lifecycle    Lifecycle
}

type Project struct {
	Issue
}

func NewProject(id int64, name string, lifecycle Lifecycle) *Project {
	project := Project{
		Issue: Issue{
			id:           id,
			name:         name,
			description:  "",
			currentState: lifecycle.startState,
			lifecycle:    lifecycle,
		},
	}
	return &project
}
