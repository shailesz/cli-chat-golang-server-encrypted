package controllers

import (
	"github.com/shailesz/cli-chat-golang-server/src/helpers"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

// SignUp creates a user.
func Signup(u, p string) int {
	hp := helpers.Sha256(p)
	services.CreateUser(u, hp)

	return 200
}

func Authenticate(u, p string) int {
	hp := helpers.Sha256(p)
	services.Login(u, hp)

	return 200
}
