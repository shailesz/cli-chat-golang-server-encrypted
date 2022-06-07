package services

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shailesz/cli-chat-golang-server/src/models"
)

// CreateUser service helps to create a new user to database.
func CreateUser(conn *pgxpool.Pool, e, u, p string) int {

	var identifier pgx.Identifier = pgx.Identifier{"users"}

	rows := [][]interface{}{
		{e, u, p},
	}

	_, err := conn.CopyFrom(
		context.TODO(),
		identifier,
		[]string{"email", "username", "password"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create user: %v\n", err)
		os.Exit(1)
	}

	return 200

}

// Login service helps validate user to database.
func Login(conn *pgxpool.Pool, u, p string) int {
	var user models.User

	const query = `SELECT username, password FROM users
	WHERE username=$1 AND password=$2`

	row := conn.QueryRow(context.TODO(), query, u, p)
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
