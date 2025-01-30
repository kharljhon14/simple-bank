package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable"
)

var testQueires *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}
	defer conn.Close(context.Background())

	testQueires = New(conn)

	// Run unit tests
	os.Exit(m.Run())
}
