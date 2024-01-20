package postgres

import (
	schema_postgres "gopt/coreapi/storage/sql/schema/postgres"
	"log"
)

func mustInitDatabase(db *postgresDatabase) {
	mustCreateTables(db)
	mustInitTablesWithData(db)
}

func mustCreateTables(db *postgresDatabase) {
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.CREATE_TABLE_LIFECYCLE_STATE); err != nil {
		log.Fatalln("ERROR :: error  creating table lifecyclestate", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.CREATE_TABLE_LIFECYCLE); err != nil {
		log.Fatalln("ERROR :: error  creating table lifecycle", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.CREATE_TABLE_STATE_TRANSITION); err != nil {
		log.Fatalln("ERROR :: error  creating table state transition", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.CREATE_TABLE_USER); err != nil {
		log.Fatalln("ERROR :: error  creating table user", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.CREATE_TABLE_PROJECT); err != nil {
		log.Fatalln("ERROR :: error  creating table project", err)
	}
}

func mustInitTablesWithData(db *postgresDatabase) {
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.INIT_LIFECYCLESTATE); err != nil {
		log.Fatalln("ERROR :: error  init table lifecyclestate", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.INIT_LIFECYCLE); err != nil {
		log.Fatalln("ERROR :: error  init table lifecycle", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.INIT_STATE_TRANSITION); err != nil {
		log.Fatalln("ERROR :: error  init table state transition", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.INIT_USERS); err != nil {
		log.Fatalln("ERROR :: error  init table users", err)
	}
}

func mustDropTables(db *postgresDatabase) {
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.DROP_TABLE_PROJECT); err != nil {
		log.Fatalln("ERROR :: error drop table project", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.DROP_TABLE_USER); err != nil {
		log.Fatalln("ERROR :: error drop table user", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.DROP_TABLE_STATE_TRANSITION); err != nil {
		log.Fatalln("ERROR :: error drop table state transition", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.DROP_TABLE_LIFECYCLE); err != nil {
		log.Fatalln("ERROR :: error drop table user", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema_postgres.DROP_TABLE_LIFECYCLE_STATE); err != nil {
		log.Fatalln("ERROR :: error drop table user", err)
	}
}
