package controllers

import (
	"github.com/shailesz/cli-chat-golang-server/src/helpers"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

// SignUp creates a user.
func Signup(e, u, p string) int {
	hp := helpers.Sha256(p)
	status := services.CreateUser(Conn, e, u, hp)

	return status
}

// Authenticate validates a user.
func Authenticate(u, p string) int {
	hp := helpers.Sha256(p)
	services.Login(Conn, u, hp)

	return 200
}
