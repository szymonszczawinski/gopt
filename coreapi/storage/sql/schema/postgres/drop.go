package postgres

const (
	DROP_TABLE_LIFECYCLE_STATE  string = "DROP TABLE IF EXISTS lifecyclestate;"
	DROP_TABLE_LIFECYCLE        string = "DROP TABLE IF EXISTS lifecycle;"
	DROP_TABLE_STATE_TRANSITION string = "DROP TABLE IF EXISTS statetransition;"
	DROP_TABLE_PROJECT          string = "DROP TABLE IF EXISTS project;"
	DROP_TABLE_USER             string = "DROP TABLE IF EXISTS users;"
)
