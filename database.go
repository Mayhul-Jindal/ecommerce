package main

import "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal.git/types"

type Database interface{
	Get(name string) (types.Book, error)
	Put() (error)
}