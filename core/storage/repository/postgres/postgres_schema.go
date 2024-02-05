package postgres

import (
	"log"
)

const (
	//------------------------------------------------------- CREATE -------------------------------------------------------

	CREATE_TABLE_LIFECYCLE_STATE  string = "CREATE TABLE IF NOT EXISTS lifecyclestate (id SERIAL PRIMARY KEY, name TEXT NOT NULL);"
	CREATE_TABLE_LIFECYCLE        string = "CREATE TABLE IF NOT EXISTS lifecycle (id SERIAL PRIMARY KEY , name TEXT NOT NULL, start_state_id INTEGER NOT NULL);"
	CREATE_TABLE_STATE_TRANSITION string = "CREATE TABLE IF NOT EXISTS statetransition (lifecycle_id INTEGER NOT NULL, from_state_id INTEGER NOT NULL, to_state_id INTEGER NOT NULL);"
	CREATE_TABLE_PROJECT          string = "CREATE TABLE IF NOT EXISTS project " +
		"(id BIGSERIAL PRIMARY KEY, created TIMESTAMP NOT NULL, updated TIMESTAMP NOT NULL, name TEXT NOT NULL, project_key TEXT NOT NULL," +
		"description TEXT, state_id INTEGER NOT  NULL, lifecycle_id INTEGER NOT NULL, created_by_id INTEGER NOT NULL, deleted boolean);"
	CREATE_TABLE_USER  string = "CREATE TABLE IF NOT EXISTS users (id BIGSERIAL PRIMARY KEY, first_name TEXT NOT NULL, last_name TEXT NOT NULL, email TEXT NOT NULL, deleted boolean);"
	CREATE_TABLE_ISSUE string = "CREATE TABLE IF NOT EXISTS issue " +
		" (id BIGSERIAL PRIMARY KEY, created TIMESTAMP NOT NULL, updated TIMESTAMP NOT NULL, name TEXT NOT NULL, item_key TEXT NOT NULL, " +
		" project_key TEXT NOT NULL, project_id INTEGER NOT NULL," +
		" description TEXT, state_id INTEGER NOT  NULL, lifecycle_id INTEGER NOT NULL, created_by_id INTEGER NOT NULL, assigned_to_id INTEGER NOT NULL, deleted boolean);"

		//------------------------------------------------------- DROP -------------------------------------------------------

	DROP_TABLE_LIFECYCLE_STATE  string = "DROP TABLE IF EXISTS lifecyclestate;"
	DROP_TABLE_LIFECYCLE        string = "DROP TABLE IF EXISTS lifecycle;"
	DROP_TABLE_STATE_TRANSITION string = "DROP TABLE IF EXISTS statetransition;"
	DROP_TABLE_PROJECT          string = "DROP TABLE IF EXISTS project;"
	DROP_TABLE_USER             string = "DROP TABLE IF EXISTS users;"
	DROP_TABLE_ISSUE            string = "DROP TABLE IF EXISTS issue;"

	//------------------------------------------------------- INIT -------------------------------------------------------

	INIT_LIFECYCLESTATE string = "INSERT INTO lifecyclestate (name) VALUES ('Draft'), ('New'), ('Open'), ('Analysis'), ('Design'), ('Development'), ('Integration')," +
		"('Verification'), ('Fixed'), ('Closed'), ('Retest'), ('Rejected'), ('Assigned');"
	INIT_LIFECYCLE        string = "INSERT INTO lifecycle (name,start_state_id) VALUES ('Project',1), ('Bug',1),('Task',1);"
	INIT_STATE_TRANSITION string = "INSERT INTO statetransition (lifecycle_id,from_state_id,to_state_id) VALUES (1,1,2), (1,2,4), (1,4,5), (1,5,6), (1,6,10);"
	INIT_USERS            string = "INSERT INTO users (first_name, last_name, email, deleted) VALUES ('Szymon','Szczawinski','szymon.szczawinski@mail.com', false);"
	INIT_PROJECTS         string = "INSERT INTO project (created, updated, name, project_key, description, state_id, lifecycle_id, created_by_id) " +
		"VALUES (NOW(),NOW(),'ICAS-1','ICAS-1','ICAS-1 description',1,1,1);"
	INIT_ISSUES string = "INSERT INTO issue (created, updated, name, item_key, project_key, project_id, description, state_id, lifecycle_id, created_by_id, assigned_to_id) " +
		"VALUES (NOW(),NOW(),'ICAS-1-1','ICAS-1-1','ICAS-1',1,' description',1,2,1,1)," +
		"(NOW(),NOW(),'ICAS-1-2','ICAS-1-2','ICAS-1',1,'ICAS-1-2 description',1,3,1,1);"
)

func mustInitDatabase(db *postgresDatabase) {
	mustCreateTables(db)
	mustInitTablesWithData(db)
}

func mustCreateTables(db *postgresDatabase) {
	if _, err := db.dbpool.Exec(db.ctx, CREATE_TABLE_LIFECYCLE_STATE); err != nil {
		log.Fatalln("ERROR :: error  creating table lifecyclestate", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, CREATE_TABLE_LIFECYCLE); err != nil {
		log.Fatalln("ERROR :: error  creating table lifecycle", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, CREATE_TABLE_STATE_TRANSITION); err != nil {
		log.Fatalln("ERROR :: error  creating table state transition", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, CREATE_TABLE_USER); err != nil {
		log.Fatalln("ERROR :: error  creating table user", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, CREATE_TABLE_PROJECT); err != nil {
		log.Fatalln("ERROR :: error  creating table project", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, CREATE_TABLE_ISSUE); err != nil {
		log.Fatalln("ERROR :: error  creating table issue", err)
	}
}

func mustInitTablesWithData(db *postgresDatabase) {
	if _, err := db.dbpool.Exec(db.ctx, INIT_LIFECYCLESTATE); err != nil {
		log.Fatalln("ERROR :: error  init table lifecyclestate", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, INIT_LIFECYCLE); err != nil {
		log.Fatalln("ERROR :: error  init table lifecycle", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, INIT_STATE_TRANSITION); err != nil {
		log.Fatalln("ERROR :: error  init table state transition", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, INIT_USERS); err != nil {
		log.Fatalln("ERROR :: error  init table users", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, INIT_PROJECTS); err != nil {
		log.Fatalln("ERROR :: error  init table projects", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, INIT_ISSUES); err != nil {
		log.Fatalln("ERROR :: error  init table issues", err)
	}
}

func mustDropTables(db *postgresDatabase) {
	if _, err := db.dbpool.Exec(db.ctx, DROP_TABLE_ISSUE); err != nil {
		log.Fatalln("ERROR :: error drop table issue", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, DROP_TABLE_PROJECT); err != nil {
		log.Fatalln("ERROR :: error drop table project", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, DROP_TABLE_USER); err != nil {
		log.Fatalln("ERROR :: error drop table user", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, DROP_TABLE_STATE_TRANSITION); err != nil {
		log.Fatalln("ERROR :: error drop table state transition", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, DROP_TABLE_LIFECYCLE_STATE); err != nil {
		log.Fatalln("ERROR :: error drop table lifecycle state", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, DROP_TABLE_LIFECYCLE); err != nil {
		log.Fatalln("ERROR :: error drop table lifecycle", err)
	}
}
