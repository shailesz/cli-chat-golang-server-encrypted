package services

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shailesz/cli-chat-golang-server/src/models"
)

// InsertToDatabase service helps to save chats to server
func InsertToDatabase(conn *pgxpool.Pool, msg models.ChatMessage) (int, error) {
	var identifier pgx.Identifier = pgx.Identifier{"chats"}

	rows := [][]interface{}{
		{msg.Username, msg.Data, msg.Timestamp},
	}

	_, err := conn.CopyFrom(
		context.TODO(),
		identifier,
		[]string{"username", "message", "timestamp"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {

		return 500, err
	}

	return 200, nil
}
