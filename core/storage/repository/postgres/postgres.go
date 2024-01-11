package postgres

import (
	"context"
	"fmt"
	"gosi/core/config"
	"gosi/coreapi/service"
	"gosi/coreapi/storage/sql/query"
	"gosi/coreapi/storage/sql/schema"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
)

type IPostgresDatabase interface {
	service.IComponent
	NewSelect(sql string, args ...any) (pgx.Rows, error)
}
type postgresDatabase struct {
	dbpool *pgxpool.Pool
	ctx    context.Context
	eg     *errgroup.Group
}

func NewPostgresSqlDatabase(eg *errgroup.Group, ctx context.Context) IPostgresDatabase {
	return &postgresDatabase{
		eg:  eg,
		ctx: ctx,
	}
}

func (db *postgresDatabase) Close() {
	db.dbpool.Close()
}

func (db *postgresDatabase) StartComponent() {
	log.Println("Starting", service.ComponentTypeSqlDatabase)
	db.dbpool = openDatabase(db.ctx)
	if config.GetSystemConfig(config.INIT_DB) == config.INIT_DB_TRUE {
		mustInitDatabase(db)
	}
	mustHealthCheck(db.dbpool, db.ctx)
}

func (db *postgresDatabase) NewSelect(sql string, args ...any) (pgx.Rows, error) {
	return db.dbpool.Query(db.ctx, sql, args...)
}

func openDatabase(ctx context.Context) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	err = dbpool.Ping(ctx)
	if err != nil {

		fmt.Fprintf(os.Stderr, "Unable to PING connection pool: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}

func mustInitDatabase(db *postgresDatabase) {
	if _, err := db.dbpool.Exec(db.ctx, schema.CREATE_TABLE_LIFECYCLE_STATE); err != nil {
		log.Fatalln("ERROR :: error  creating table lifecyclestate", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.INIT_LIFECYCLESTATE); err != nil {
		log.Fatalln("ERROR :: error  init table lifecyclestate", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.CREATE_TABLE_LIFECYCLE); err != nil {
		log.Fatalln("ERROR :: error  creating table lifecycle", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.INIT_LIFECYCLE); err != nil {
		log.Fatalln("ERROR :: error  init table lifecycle", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.CREATE_TABLE_STATE_TRANSITION); err != nil {
		log.Fatalln("ERROR :: error  creating table state transition", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.INIT_STATE_TRANSITION); err != nil {
		log.Fatalln("ERROR :: error  init table state transition", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.CREATE_TABLE_USER); err != nil {
		log.Fatalln("ERROR :: error  creating table user", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.INIT_USERS); err != nil {
		log.Fatalln("ERROR :: error  init table users", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.CREATE_TABLE_PROJECT); err != nil {
		log.Fatalln("ERROR :: error  creating table project", err)
	}
	if _, err := db.dbpool.Exec(db.ctx, schema.INIT_PROJECT); err != nil {
		log.Fatalln("ERROR :: error  init table projects", err)
	}
}

func mustHealthCheck(pool *pgxpool.Pool, ctx context.Context) {
	rows, err := pool.Query(ctx, query.PROJECT_SELECT_ALL)
	if err != nil {
		log.Panicln("ERROR :: DB healthceck", err)
	}
	defer rows.Close()
}
