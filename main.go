// TODO
/*
- simple message in run
*/
package main

import (
	"log"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	razorpay "github.com/razorpay/razorpay-go"
)

func main() {
	// tools, database etc
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config")
	}

	database := NewPostgresStore()
	tokenMaker := token.NewPasetoMaker(config.TOKEN_SYMMETRIC_KEY)
	razorPayclient := razorpay.NewClient(config.RAZORPAY_KEY_ID, config.RAZORPAY_KEY_SECRET)

	// this is the auth serive invocation
	authService := NewAuthManager(config, tokenMaker, database)

	// this is the book serive invocation
	bookService := NewBookManager(database, razorPayclient)

	server := NewAPIServer(authService, bookService, tokenMaker)
	server.Run()
}
