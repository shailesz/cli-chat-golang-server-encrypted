package controllers

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/shailesz/cli-chat-golang-server/src/helpers"
	"github.com/shailesz/cli-chat-golang-server/src/models"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

// SignUp creates a user.
func Signup(user models.User) int {
	hp := helpers.Sha256(user.Password)

	user = models.User{
		Email:     user.Email,
		Username:  user.Username,
		Password:  hp,
		PublicKey: user.PublicKey,
	}

	status := services.CreateUser(Conn, user)

	return status
}

// Authenticate validates a user.
func Authenticate(u, p string) int {
	_, statusCode := services.Login(Conn, u, p)

	return statusCode
}

func getPublicKeyForUser(conn socketio.Conn, username string) {
	key := services.GetPublicKey(Conn, username)

	conn.Emit("getPublicKey", key)
}
