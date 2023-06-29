package sqlite

import (
	"database/sql"
	"gosi/core/storage/sql/command"
	"gosi/core/storage/sql/schema"
	"gosi/issues/domain"
	"log"
	"time"
)

func initaliseDatabase(database *sql.DB) {
	log.Println("initaliseDatabase")
	createSchema(database)
	initialiseLifecycleStates(database)
	initialiseLifecycles(database)
	initialiseStateTransitions(database)
	initialiseProjects(database)
}

func createSchema(db *sql.DB) {
	createTable(db, schema.CREATE_TABLE_LIFECYCLE_STATE)
	createTable(db, schema.CREATE_TABLE_LIFECYCLE)
	createTable(db, schema.CREATE_TABLE_STATE_TRANSITION)
	createTable(db, schema.CREATE_TABLE_PROJECT)
}

func initialiseLifecycleStates(db *sql.DB) {
	log.Println("Initialise LifecycleStates")
	tableIsEmpty := isTableEmpty(db, "select count(id) as statesnumber from lifecyclestate;")

	if tableIsEmpty {
		insertStatement, _ := db.Prepare(command.INSERT_LIFECYCLE_STATE)
		defer insertStatement.Close()
		insertStatement.Exec(domain.LIFECYCLE_STATE_DRAFT)
		insertStatement.Exec(domain.LIFECYCLE_STATE_NEW)
		insertStatement.Exec(domain.LIFECYCLE_STATE_OPEN)
		insertStatement.Exec(domain.LIFECYCLE_STATE_ANALISYS)
		insertStatement.Exec(domain.LIFECYCLE_STATE_DESIGN)
		insertStatement.Exec(domain.LIFECYCLE_STATE_DEVELOPMENT)
		insertStatement.Exec(domain.LIFECYCLE_STATE_INTEGRATION)
		insertStatement.Exec(domain.LIFECYCLE_STATE_VERIFICATION)
		insertStatement.Exec(domain.LIFECYCLE_STATE_FIXED)
		insertStatement.Exec(domain.LIFECYCLE_STATE_CLOSED)
		insertStatement.Exec(domain.LIFECYCLE_STATE_RETEST)
		insertStatement.Exec(domain.LIFECYCLE_STATE_REJECTED)
	}

}

func initialiseLifecycles(db *sql.DB) {
	log.Println("Initialise Lifecycles")
	tableIsEmpty := isTableEmpty(db, "select count(id) as lcnumber from lifecycle;")
	if tableIsEmpty {
		insertStatement, _ := db.Prepare(command.INSERT_LIFECYCLE)
		defer insertStatement.Close()
		insertStatement.Exec("Project", 1)
	}

}

func initialiseStateTransitions(db *sql.DB) {
	log.Println("Initialise StateTransitions")
	tableIsEmpty := isTableEmpty(db, "select count(lifecycleid) as transitionsnumber from statetransition;")

	if tableIsEmpty {

		insertStatement, _ := db.Prepare(command.INSERT_STATE_TRANSITION)
		defer insertStatement.Close()
		//Project :: DRAFT -> NEW -> ANALISYS -> DESIGN -> DEV -> CLOSED
		insertStatement.Exec(1, 1, 2)
		insertStatement.Exec(1, 2, 4)
		insertStatement.Exec(1, 4, 5)
		insertStatement.Exec(1, 5, 6)
		insertStatement.Exec(1, 6, 10)
	}

}

func initialiseProjects(db *sql.DB) {
	log.Println("Initialise projects")
	tableIsEmpty := isTableEmpty(db, "select count(id) as projectid from project;")

	if tableIsEmpty {
		//(id, created, updated, name, itemkey, intemnumber,description, stateid, lifecycleid, createdbyid)
		insertStatement, err := db.Prepare(command.INSERT_PROJECT)
		if err != nil {
			log.Println(err.Error())
		}
		defer insertStatement.Close()
		insertStatement.Exec(time.Now(), time.Now(), "COMSOS", "COMSOS", 1, "Cosmos Project", 1, 1, 1)
	}
}
func createTable(db *sql.DB, command string) error {
	_, err := db.Exec(command)
	if err != nil {
		log.Println("Error in creating table")
		return err
	}
	log.Println("Successfully created table!")
	return nil

}
func isTableEmpty(db *sql.DB, query string) bool {
	rows, _ := db.Query(query)
	defer rows.Close()
	var rowNumber int
	for rows.Next() {
		rows.Scan(&rowNumber)
		break
	}
	log.Println("Rows number: ", rowNumber)
	return rowNumber == 0
}
