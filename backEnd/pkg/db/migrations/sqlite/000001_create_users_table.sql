-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    nickname TEXT UNIQUE,
    about_me TEXT,
    profile_visibility TEXT NOT NULL CHECK (profile_visibility IN ('public', 'private')),
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    birthday TEXT NOT NULL,
    avatar TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down

DROP TABLE IF EXISTS users;