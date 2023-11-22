package sql

import (
	"context"
	"fmt"
	"gosi/coreapi/service"
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
