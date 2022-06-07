package main

import (
	"log"
	"net/http"

	"github.com/shailesz/cli-chat-golang-server/src/controllers"
)

func main() {

	// init app
	controllers.InitApp()

	// serve socket.io server | handle websockets
	go func() {
		if err := controllers.Server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	// defer close server
	defer controllers.Server.Close()

	// handle http for server on defined route
	http.Handle("/socket.io/", controllers.Server)

	// start http server for init connection
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
