package postgres

const (
	CREATE_TABLE_LIFECYCLE_STATE  string = "CREATE TABLE IF NOT EXISTS lifecyclestate (id SERIAL PRIMARY KEY, name TEXT NOT NULL);"
	CREATE_TABLE_LIFECYCLE        string = "CREATE TABLE IF NOT EXISTS lifecycle (id SERIAL PRIMARY KEY , name TEXT NOT NULL, start_state_id INTEGER NOT NULL);"
	CREATE_TABLE_STATE_TRANSITION string = "CREATE TABLE IF NOT EXISTS statetransition (lifecycle_id INTEGER NOT NULL, from_state_id INTEGER NOT NULL, to_state_id INTEGER NOT NULL);"
	CREATE_TABLE_PROJECT          string = "CREATE TABLE IF NOT EXISTS project " +
		"(id BIGSERIAL PRIMARY KEY, created TIMESTAMP NOT NULL, updated TIMESTAMP NOT NULL, name TEXT NOT NULL, project_key TEXT NOT NULL," +
		"description TEXT, state_id INTEGER NOT  NULL, lifecycle_id INTEGER NOT NULL, created_by_id INTEGER NOT NULL, deleted boolean);"
	CREATE_TABLE_USER string = "CREATE TABLE IF NOT EXISTS users (id BIGSERIAL PRIMARY KEY, first_name TEXT NOT NULL, last_name TEXT NOT NULL, email TEXT NOT NULL, deleted boolean);"
)