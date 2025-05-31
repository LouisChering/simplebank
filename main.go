package main

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/louischering/simplebank/api"
	db "github.com/louischering/simplebank/db/sqlc"
	"github.com/louischering/simplebank/util"
)

//go:embed views/*
var templateFS embed.FS

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, config.DBConnectionString)
	if err != nil {
		log.Fatal("unable to connect to db.", err)
	}
	defer dbpool.Close()

	store := db.NewStore(dbpool)
	server := api.NewServer(store, &templateFS)

	if err = server.Start(config.ServerAddress); err != nil {
		fmt.Printf("error: %v", err)
	}
}
