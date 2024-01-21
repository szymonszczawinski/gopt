package postgres

import (
	"context"
	"fmt"
	"gopt/core/config"
	"gopt/coreapi/service"
	"gopt/coreapi/storage/sql/query"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
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

func mustHealthCheck(pool *pgxpool.Pool, ctx context.Context) {
	rows, err := pool.Query(ctx, query.PROJECT_SELECT_ALL)
	if err != nil {
		log.Panicln("ERROR :: DB healthceck", err)
	}
	defer rows.Close()
}
