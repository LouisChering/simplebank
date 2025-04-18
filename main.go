package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/louischering/simplebank/api"
	db "github.com/louischering/simplebank/db/sqlc"
)

const (
	connString = "host=localhost port=5432 password=mysecretpassword user=postgres dbname=simple_bank sslmode=disable"
)

func main() {
	var err error

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Fatal("unable to connect to db.", err)
	}
	defer dbpool.Close()

	store := db.NewStore(dbpool)
	server := api.NewServer(store)

	if err = server.Start("0.0.0.0:8080"); err != nil {
		fmt.Printf("error: %v", err)
	}
}
