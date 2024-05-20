CREATE TABLE IF NOT EXISTS users(
    UserId TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);