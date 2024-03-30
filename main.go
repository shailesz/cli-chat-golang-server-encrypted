package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

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

	// wait for a sigint terminate with ctrl+c
	c := make(chan os.Signal, 1)
	go func(c chan os.Signal) {
		signal.Notify(c, os.Interrupt)
	}(c)

	go func(chan os.Signal) {
		for range c {
			log.Println("Shutting down server...")
			controllers.Server.Close()
			os.Exit(0)
		}
	}(c)

	// start http server for init connection
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
