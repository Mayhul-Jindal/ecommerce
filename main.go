// TODO
/*
- simple message in run
*/
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/authService"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/bookService"
	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/email"
	gatewayservice "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/gatewayService"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/worker"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	razorpay "github.com/razorpay/razorpay-go"
)

func main() {
	// config tool
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config: %v\n", err)
		os.Exit(1)
	}

	// validator tool
	validator := validator.New()

	// database 
	dbPool, err := pgxpool.New(context.Background(), config.DB_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbPool.Close()
	database := database.NewPostgresStore(dbPool)

	// token maker
	tokenMaker := token.NewPasetoMaker(config.TOKEN_SYMMETRIC_KEY)

	// emailer
	emailer := email.NewGmailSender(config.EMAIL_SENDER_NAME, config.EMAIL_SENDER_ADDRESS, config.EMAIL_SENDER_PASSWORD)

	// racor pay client
	razorPayClient := razorpay.NewClient(config.RAZORPAY_KEY_ID, config.RAZORPAY_KEY_SECRET)

	// worker service
	worker := worker.NewWorker(database, emailer)

	// this is the auth serive invocation
	authService := authService.NewManager(config, tokenMaker, database, worker)

	// this is the book serive invocation
	bookService := bookService.NewManager(database, razorPayClient)

	server := gatewayservice.NewAPIServer(config.SERVER_PORT, authService, bookService, tokenMaker, validator)
	server.Run()
}
