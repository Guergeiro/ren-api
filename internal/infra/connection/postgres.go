package connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PostgresConn *pgxpool.Pool

func init() {
	url, ok := os.LookupEnv("POSTGRES_URL")
	if ok == false {
		panic("No Postgres url environment variable")
	}
	conn, err := pgxpool.New(context.Background(), url)
	if err != nil {
		panic("Can't establish Postgres connection")
	}
	PostgresConn = conn
}
