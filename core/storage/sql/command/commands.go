package command

const (
	INSERT_LIFECYCLE_STATE  string = "INSERT INTO lifecyclestate (id, name) VALUES (NULL,?);"
	INSERT_LIFECYCLE        string = "INSERT INTO lifecycle (id, name,startstateid) VALUES (NULL,?,?);"
	INSERT_STATE_TRANSITION string = "INSERT INTO statetransition (lifecycleid, fromstateid, tostateid) VALUES (?,?,?);"
	INSERT_PROJECT          string = "INSERT INTO project (id, created, updated, name, itemkey, itemnumber," +
		"description, stateid, lifecycleid, createdbyid) VALUES(NULL,?,?,?,?,?,?,?,?,?);"
)
