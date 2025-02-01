package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kharljhon14/simple-bank/cmd/api"
	db "github.com/kharljhon14/simple-bank/db/sqlc"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), os.Getenv("DSN"))
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(os.Getenv("ADDRESS"))
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
