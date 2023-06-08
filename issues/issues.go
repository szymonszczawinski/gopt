package issues

type Issue struct {
	Id           int64          `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	CurrentState LifecycleState `json:"state"`
	Lifecycle    Lifecycle      `json:"lifecycle"`
}

type Project struct {
	Issue
}

func NewProject(id int64, name string, lifecycle Lifecycle) *Project {
	project := Project{
		Issue: Issue{
			Id:           id,
			Name:         name,
			Description:  "",
			CurrentState: lifecycle.startState,
			Lifecycle:    lifecycle,
		},
	}
	return &project
}
