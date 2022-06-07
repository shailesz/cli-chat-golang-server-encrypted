package controllers

import (
	"log"

	"github.com/shailesz/cli-chat-golang-server/src/models"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

// SaveChat saves chat to database.
func SaveChat(msg models.ChatMessage) {
	_, err := services.InsertToDatabase(Conn, msg)

	if err != nil {
		log.Println("could not save chat to server.", err)
	}
}
