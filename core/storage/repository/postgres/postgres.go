package postgres

import (
	"context"
	"fmt"
	"gopt/core/config"
	"gopt/coreapi/service"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"golang.org/x/sync/errgroup"
)

const (
	PROJECT_SELECT_ALL string = "SELECT id, created, updated, name, project_key, description, state_id," +
		"lifecycle_id, created_by_id FROM project;"

	LIFECYCLE_SELECT_ALL       string = "SELECT id, name, start_state_id FROM lifecycle;"
	LIFECYCLE_STATE_SELECT_ALL string = "SELECT id, name FROM lifecyclestate;"
)

type IPostgresDatabase interface {
	service.IComponent
	NewSelect(sql string, args ...any) (pgx.Rows, error)
	NewSelectOne(sql string, args any) pgx.Row
	NewInsert(sql string, args ...any) (pgconn.CommandTag, error)
	NewInsertReturninId(sql string, args any) (int, error)
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
		mustDropTables(db)
		mustInitDatabase(db)
	}
	mustHealthCheck(db.dbpool, db.ctx)
}

func (db *postgresDatabase) NewSelect(sql string, args ...any) (pgx.Rows, error) {
	return db.dbpool.Query(db.ctx, sql, args...)
}

func (db *postgresDatabase) NewSelectOne(sql string, args any) pgx.Row {
	return db.dbpool.QueryRow(db.ctx, sql, args)
}

func (db *postgresDatabase) NewInsert(sql string, args ...any) (pgconn.CommandTag, error) {
	return db.dbpool.Exec(db.ctx, sql, args...)
}

func (db *postgresDatabase) NewInsertReturninId(sql string, args any) (int, error) {
	var id int
	err := db.dbpool.QueryRow(db.ctx, sql, args).Scan(&id)
	return id, err
}

type myLogger struct{}

func (ll myLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	log.Println("PGX", "level=", level, "msg=", msg, "args=", data)
}

func openDatabase(ctx context.Context) *pgxpool.Pool {
	tracer := &tracelog.TraceLog{
		Logger:   myLogger{},
		LogLevel: tracelog.LogLevelTrace,
	}

	dbConfig, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connString: %v\n", err)
		os.Exit(1)
	}

	dbConfig.ConnConfig.Tracer = tracer
	dbpool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	// dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
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

func mustHealthCheck(pool *pgxpool.Pool, ctx context.Context) {
	rows, err := pool.Query(ctx, PROJECT_SELECT_ALL)
	if err != nil {
		log.Panicln("ERROR :: DB healthceck", err)
	}
	defer rows.Close()
}
