package services

import (
	"context"
	"fmt"
	"log"
	"os"

	socketio "github.com/googollee/go-socket.io"
	"github.com/jackc/pgx/v4/pgxpool"
)

// InitWebsocket helps to create a new websocket server initialization.
func InitWebsocket() *socketio.Server {
	server := socketio.NewServer(nil)

	return server
}

func InitDatabaseConnection() *pgxpool.Pool {
	// hardcoded db url
	databaseUrl := "postgres://shailesz:password@localhost:5432/cli-chat-golang"

	// this returns connection pool
	conn, err := pgxpool.Connect(context.Background(), databaseUrl)

	// handle error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	log.Println("Database connection successful.")

	return conn
}
