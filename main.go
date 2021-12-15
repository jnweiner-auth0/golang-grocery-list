package main

// TODO: uncomment below once you've generated your rpc files and you're ready to start the server

import (
	"fmt"
	config "grocery-list/config"
	groceryProto "grocery-list/rpc/grocery"
	"grocery-list/server"
	"net/http"
)

func startServer() {
	fmt.Println("Starting server")

	server := &server.Server{}
	handler := groceryProto.NewGroceryServiceServer(server)

	fmt.Printf("Server listening on port: %v\n", config.Port)

	listener := fmt.Sprintf(":%v", config.Port)
	http.ListenAndServe(listener, handler)
}

func main() {
	fmt.Println("Hello world")
	config.ConnectToPostgres()
	startServer()
}
