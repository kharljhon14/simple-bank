package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable"
)

var testStore Store

func TestMain(m *testing.M) {
	var err error

	testDBPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}

	testStore = *NewStore(testDBPool)

	// Run unit tests
	os.Exit(m.Run())
}
