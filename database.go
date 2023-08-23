// TODO
/*
- postgres connection
- creating test for these things
- schema migration
*/

package main

import "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"

type Database interface{
	Get(name string) (types.Book, error)
	Put() (error)
}