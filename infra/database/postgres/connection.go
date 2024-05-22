package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	POSTGRES_USER     = "POSTGRES_USER"
	POSTGRES_PASSWORD = "POSTGRES_PASSWORD"
	POSTGRES_DB       = "POSTGRES_DB"
	POSTGRES_HOST     = "POSTGRES_HOST"
	POSTGRES_PORT     = "POSTGRES_PORT"
)

var (
	username    = os.Getenv(POSTGRES_USER)
	password    = os.Getenv(POSTGRES_PASSWORD)
	host        = os.Getenv(POSTGRES_HOST)
	db          = os.Getenv(POSTGRES_DB)
	port        = os.Getenv(POSTGRES_PORT)
	databaseURL = fmt.Sprintf("://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, db)
)

type PoolInterface interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(
		ctx context.Context,
		sql string,
		args []interface{},
		scans []interface{},
		f func(pgx.QueryFuncRow) error,
	) (pgconn.CommandTag, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

func GetConnection(context context.Context) *pgxpool.Pool {
	conn, err := pgxpool.Connect(context, "postgres"+databaseURL)

	if err != nil {
		panic(err)
	}

	return conn
}

// RunMigrations run scripts on path database/migrations
func RunMigrations() {
	m, err := migrate.New("file://infra/database/postgres/migrations", "pgx"+databaseURL)
	if err != nil {
		log.Println("Erro run migration 1: ", err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
