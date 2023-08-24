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
	"os"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storer interface {
	postgres.Querier
}

func NewPostgresStore() Storer {
	dbPool, err := pgxpool.New(context.Background(), "postgresql://admin:admin@localhost:5432/book-store?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	
	queries := postgres.New(dbPool)
	return queries
}
