package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	connString = "host=localhost port=5432 password=mysecretpassword user=postgres dbname=simple_bank sslmode=disable"
)

var testQueries *Queries
var dbpool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	ctx := context.Background()
	dbpool, err = pgxpool.New(ctx, connString)
	// connection, err = pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatal("unable to connect to db.", err)
	}
	// defer connection.Close(ctx)
	testQueries = New(dbpool)
	os.Exit(m.Run())
}
