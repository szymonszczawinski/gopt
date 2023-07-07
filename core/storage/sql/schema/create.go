package schema

const (
	CREATE_TABLE_LIFECYCLE_STATE  string = "CREATE TABLE IF NOT EXISTS lifecyclestate (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);"
	CREATE_TABLE_LIFECYCLE        string = "CREATE TABLE IF NOT EXISTS lifecycle (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, start_state_id INTEGER);"
	CREATE_TABLE_STATE_TRANSITION string = "CREATE TABLE IF NOT EXISTS statetransition (lifecycle_id INTEGER, from_state_id INTEGER, to_state_id INTEGER );"
	CREATE_TABLE_PROJECT          string = "CREATE TABLE IF NOT EXISTS project " +
		"(id INTEGER PRIMARY KEY AUTOINCREMENT, created DATETIME , updated DATETIME, name TEXT, item_key TEXT, item_number INTEGER," +
		"description TEXT, state_id INTEGER, lifecycle_id INTEGER, created_by_id INTEGER);"
)
