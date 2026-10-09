-- +migrate Up

CREATE TABLE IF NOT EXISTS Sessions (
    user_id INTEGER NOT NULL,
    session_id TEXT NOT NULL UNIQUE,
    expired_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);

-- +migrate Down

DROP TABLE IF EXISTS Sessions;