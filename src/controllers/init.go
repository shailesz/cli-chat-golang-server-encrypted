package controllers

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

// connection variables
var Conn *pgxpool.Pool
var Server *socketio.Server

// InitApp initializes connection variables.
func InitApp() {
	Conn = services.InitDatabaseConnection()
	Server = services.InitWebsocket()

	Server.OnConnect("/", OnConnectHandler)
	Server.OnEvent("/", "auth", LoginHandler)
	Server.OnEvent("/", "signup", SignupHandler)
	Server.OnEvent("/", "chat", ChatHandler)
	Server.OnError("/", ErrorHandler)

	log.Println("Websocket server setup!")
}
