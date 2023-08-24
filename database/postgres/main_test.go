package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = ""
	dbUrl    = "postgresql://admin:admin@localhost:5432/book-store?sslmode=disable"
)

// This struct will be used by other tests. This struct has methods
var testQueries *Queries

func TestMain(m *testing.M) {
	// This dbPool is implementing DBTX interface or else I would not be able to put this into New()
	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	testQueries = New(dbPool)
	os.Exit(m.Run())
}
