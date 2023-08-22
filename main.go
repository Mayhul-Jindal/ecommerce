package main

import "fmt"

func main() {
	fmt.Println("testing my service")
	
	bussinessLogic := NewBookManager()
	blWithLogging := NewLoggingService(bussinessLogic)
	server := NewAPIServer(":3000", blWithLogging)
	server.Run()
}

