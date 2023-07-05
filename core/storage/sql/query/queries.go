package query

const (
	PROJECT_SELECT_ALL         string = "SELECT id, created, updated, name, itemkey, itemnumber, description, stateid, lifecycleid, createdbyid FROM project;"
	LIFECYCLE_SELECT_ALL       string = "SELECT id, name, startstateid FROM lifecycle;"
	LIFECYCLE_STATE_SELECT_ALL string = "SELECT id, name FROM lifecyclestate;"
)
