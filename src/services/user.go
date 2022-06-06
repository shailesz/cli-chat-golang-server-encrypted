package services

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/shailesz/cli-chat-golang-server/src/models"
)

func CreateUser(u, p string) int {

	var identifier pgx.Identifier = pgx.Identifier{"users"}

	rows := [][]interface{}{
		{u, p},
	}

	_, err := Conn.CopyFrom(
		context.TODO(),
		identifier,
		[]string{"username", "password"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create user: %v\n", err)
		os.Exit(1)
	}

	return 200

}

func Login(u, p string) int {
	var user models.User

	const query = `SELECT username, password FROM users
	WHERE username=$1 AND password=$2`

	row := Conn.QueryRow(context.TODO(), query, u, p)
	if row != nil {
		err := row.Scan(&user.Username, &user.Password)

		if err != nil {
			return 404
		}
	} else {
		return 404
	}

	return 200
}
