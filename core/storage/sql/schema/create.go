package schema

const (
	CREATE_TABLE_LIFECYCLE_STATE  string = "CREATE TABLE IF NOT EXISTS lifecyclestate (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT);"
	CREATE_TABLE_LIFECYCLE        string = "CREATE TABLE IF NOT EXISTS lifecycle (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, startstateid INTEGER);"
	CREATE_TABLE_STATE_TRANSITION string = "CREATE TABLE IF NOT EXISTS statetransition (lifecycleid INTEGER, fromstateid INTEGER, tostateid INTEGER );"
	CREATE_TABLE_PROJECT          string = "CREATE TABLE IF NOT EXISTS project " +
		"(id INTEGER PRIMARY KEY AUTOINCREMENT, created DATETIME , updated DATETIME, name TEXT, itemkey TEXT, itemnumber INTEGER," +
		"description TEXT, stateid INTEGER, lifecycleid INTEGER, createdbyid INTEGER);"
)
