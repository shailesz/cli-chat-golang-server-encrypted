package handlers

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/shailesz/cli-chat-golang-server/src/controllers"
	"github.com/shailesz/cli-chat-golang-server/src/models"
)

var Server *socketio.Server = InitHandlers()

func InitHandlers() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", OnConnectHandler)
	server.OnEvent("/", "auth", LoginHandler)
	server.OnEvent("/", "signup", SignupHandler)
	server.OnEvent("/", "chat", func(s socketio.Conn, msg models.ChatMessage) {
		server.BroadcastToRoom("/", "chatroom", "message", msg)
	})
	server.OnError("/", ErrorHandler)

	log.Println("Websocket server setup!")

	return server
}

func OnConnectHandler(s socketio.Conn) error {
	s.SetContext("")
	log.Println("connected:", s.ID())
	s.Join("chatroom")
	return nil
}

func LoginHandler(s socketio.Conn, user models.User) {
	status := controllers.Authenticate(user.Username, user.Password)

	if status == 200 {
		res := models.AuthMessage{Status: 200, Data: user}
		s.Emit("auth", res)
	} else {
		res := models.AuthMessage{Status: 404, Data: user}
		s.Emit("auth", res)
	}
}

func SignupHandler(s socketio.Conn, user models.User) {
	status := controllers.Signup(user.Username, user.Password)

	if status == 200 {
		res := models.AuthMessage{Status: 200, Data: user}
		s.Emit("signup", res)
	} else {
		res := models.AuthMessage{Status: 404, Data: user}
		s.Emit("signup", res)
	}
}

func ErrorHandler(s socketio.Conn, e error) {
	log.Panicln("meet error:", e)
}
