package main

import "fmt"

func main() {
	fmt.Println("testing my service")
	
	database := NewPostgresStore()
	// bussinessLogic := NewBookManager(database)
	// blWithLogging := NewLoggingService(bussinessLogic)
	server := NewAPIServer(":3000", database)
	server.Run()
}

