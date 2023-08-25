package sqlc_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This struct will be used by other tests. This struct has methods
var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	// This dbPool is implementing DBTX interface or else I would not be able to put this into New()
	dbPool, err := pgxpool.New(context.Background(), config.DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	testQueries = sqlc.New(dbPool)
	os.Exit(m.Run())
}
