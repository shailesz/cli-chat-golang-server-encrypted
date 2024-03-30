package controllers

import (
	"database/sql"
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

// connection variables
var Conn *sql.DB
var Server *socketio.Server

// InitApp initializes connection variables.
func InitApp() {
	Conn = services.InitDatabaseConnection()
	Server = services.InitWebsocket()

	Server.OnConnect("/", OnConnectHandler)
	Server.OnEvent("/", "auth", LoginHandler)
	Server.OnEvent("/", "signup", SignupHandler)
	Server.OnEvent("/", "chat", ChatHandler)
	Server.OnEvent("/", "getPublicKey", getPublicKeyForUser)
	Server.OnError("/", ErrorHandler)

	log.Println("Websocket server setup!")
}
