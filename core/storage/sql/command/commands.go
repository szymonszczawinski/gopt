package command

const (
	INSERT_LIFECYCLE_STATE  string = "INSERT INTO lifecyclestate (id, name) VALUES (NULL,?);"
	INSERT_LIFECYCLE        string = "INSERT INTO lifecycle (id, name,start_state_id) VALUES (NULL,?,?);"
	INSERT_STATE_TRANSITION string = "INSERT INTO statetransition (lifecycleid, fromstateid, to_state_id) VALUES (?,?,?);"
	INSERT_PROJECT          string = "INSERT INTO project (id, created, updated, name, item_key, item_number," +
		"description, state_id, lifecycle_id, created_by_id) VALUES(NULL,?,?,?,?,?,?,?,?,?);"
)
