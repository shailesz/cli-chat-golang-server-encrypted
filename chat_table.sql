CREATE TABLE users
(
    user_id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE chats
(
    chat_id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    message TEXT,
    timestamp BIGINT NOT NULL,
    FOREIGN KEY(username) 
    REFERENCES users(username)
);