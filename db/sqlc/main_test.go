package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/louischering/simplebank/util"
)

var testQueries *Queries
var dbpool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	ctx := context.Background()
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cant load config")
	}
	dbpool, err = pgxpool.New(ctx, config.DBConnectionString)
	// connection, err = pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatal("unable to connect to db.", err)
	}
	// defer connection.Close(ctx)
	testQueries = New(dbpool)
	os.Exit(m.Run())
}
