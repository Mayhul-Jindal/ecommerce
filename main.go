// TODO
/*
- simple message in run
*/
package main

func main() {
	database := NewPostgresStore()

	// this is the auth serive invocation
	authService := NewAuthManager(database)

	// this is the book serive invocation
	bookService := NewLoggingService(NewBookManager(database))

	server := NewAPIServer(authService, bookService)
	server.Run()
}
