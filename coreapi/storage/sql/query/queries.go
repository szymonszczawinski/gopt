package query

const (
	PROJECT_SELECT_ALL string = "SELECT id, created, updated, name, project_key, description, state_id," +
		"lifecycle_id, created_by_id FROM project;"

	LIFECYCLE_SELECT_ALL       string = "SELECT id, name, start_state_id FROM lifecycle;"
	LIFECYCLE_STATE_SELECT_ALL string = "SELECT id, name FROM lifecyclestate;"
)
