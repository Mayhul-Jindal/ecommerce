// TODO
/*
- postgres connection
- creating test for these things
- schema migration
*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storer interface {
	database.Querier
}

func NewPostgresStore() Storer {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	dbPool, err := pgxpool.New(context.Background(), config.DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	// defer dbPool.Close()

	queries := database.New(dbPool)
	return queries
}
