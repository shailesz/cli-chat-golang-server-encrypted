package controllers

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/shailesz/cli-chat-golang-server/src/models"
)

// ChatHandler handles outgoing messages to connected users.
func ChatHandler(s socketio.Conn, msg models.ChatMessage) {
	Server.BroadcastToRoom("/", "chatroom", "message", msg)
	SaveChat(msg)
}

// OnConnectHandler handles client on connect.
func OnConnectHandler(s socketio.Conn) error {
	s.SetContext("")
	log.Println("connected:", s.ID())
	s.Join("chatroom")
	return nil
}

// LoginHandler handles login/authentication messages.
func LoginHandler(s socketio.Conn, user models.User) {
	status := Authenticate(user.Username, user.Password)

	if status == 200 {
		res := models.AuthMessage{Status: 200, Data: user}
		s.Emit("auth", res)
	} else {
		res := models.AuthMessage{Status: 404, Data: user}
		s.Emit("auth", res)
	}
}

// SignupHandler handles signup messages.
func SignupHandler(s socketio.Conn, user models.User) {
	status := Signup(user)

	res := models.AuthMessage{Status: status, Data: user}
	s.Emit("signup", res)
}

// ErrorHandler handles error messages.
func ErrorHandler(s socketio.Conn, e error) {
	log.Panicln("meet error:", e)
}
