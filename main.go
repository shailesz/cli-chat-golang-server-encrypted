package main

import (
	"log"
	"net/http"

	"github.com/shailesz/cli-chat-golang-server/src/handlers"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

func main() {

	services.InitApp()

	// serve socket.io server | handle websockets
	go func() {
		if err := handlers.Server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	// defer close server
	defer handlers.Server.Close()

	// handle http for server on defined route
	http.Handle("/socket.io/", handlers.Server)

	// start http server for init connection
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
