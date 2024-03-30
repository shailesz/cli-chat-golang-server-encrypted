package services

import (
	"context"
	"fmt"
	"os"

	"database/sql"

	"github.com/shailesz/cli-chat-golang-server/src/helpers"
	"github.com/shailesz/cli-chat-golang-server/src/models"
)

// CreateUser service helps to create a new user in the database.
func CreateUser(db *sql.DB, user models.User) int {
	// SQL INSERT statement to add a new user
	query := `INSERT INTO users (email, username, password, public_key) VALUES (?, ?, ?, ?)`

	// Executing the INSERT statement
	_, err := db.ExecContext(context.Background(), query, user.Email, user.Username, user.Password, user.PublicKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating user: %v\n", err)
		return 409
	}

	return 200
}

func Login(db *sql.DB, username, password string) (models.User, int) {
	var user models.User

	// Include the public_key in the SELECT query
	const query = `SELECT username, password, public_key FROM users WHERE username=? AND password=?`

	// Query the database for the user
	row := db.QueryRowContext(context.Background(), query, username, helpers.Sha256(password)) // Assuming password is hashed
	err := row.Scan(&user.Username, &user.Password, &user.PublicKey)

	if err != nil {
		// If an error occurs (e.g., no rows found), return an empty user and 404
		fmt.Println(err.Error())
		return models.User{}, 404
	}

	// Return the user and 200 on successful login
	return user, 200
}

func GetPublicKey(db *sql.DB, username string) string {
	var user models.User

	const query = `SELECT public_key FROM users WHERE username=?`

	row := db.QueryRowContext(context.Background(), query, username)
	err := row.Scan(&user.PublicKey)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return user.PublicKey
}
