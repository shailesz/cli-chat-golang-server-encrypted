package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	socketio "github.com/googollee/go-socket.io"
	_ "github.com/mattn/go-sqlite3"
)

// InitWebsocket helps to create a new websocket server initialization.
func InitWebsocket() *socketio.Server {
	server := socketio.NewServer(nil)

	return server
}

// func InitDatabaseConnection() *pgxpool.Pool {
// 	// hardcoded db url
// 	databaseUrl := "postgres://shailesz:password@localhost:5432/cli-chat-golang"

// 	// this returns connection pool
// 	conn, err := pgxpool.Connect(context.Background(), databaseUrl)

// 	// handle error
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	log.Println("Database connection successful.")

// 	return conn
// }

func InitDatabaseConnection() *sql.DB {
	databaseFile := "chat-golang.db"

	// Open a database connection, SQLite will create the file if it does not exist.
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open database file: %v\n", err)
		os.Exit(1)
	}

	// You might want to ping the database to ensure the connection is established.
	err = db.PingContext(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Println("Database connection successful.")

	InitSeedDB(db)

	return db
}

func InitSeedDB(conn *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users
				(
					user_id INTEGER PRIMARY KEY,
					email TEXT UNIQUE NOT NULL,
					username TEXT UNIQUE NOT NULL,
					password TEXT NOT NULL,
					public_key TEXT NOT NULL
				);
				`

	// execute query
	_, err := conn.Exec(query)
	if err != nil {
		log.Println("Table already exists.")
	}

	query = `CREATE TABLE chats
			(
				chat_id INTEGER PRIMARY KEY,
				username TEXT NOT NULL,
				message TEXT,
				timestamp INTEGER NOT NULL,
				FOREIGN KEY(username) REFERENCES users(username)
			);
			`

	// execute query
	_, err = conn.Exec(query)
	if err != nil {
		log.Println("Table already exists.")
	}

}
