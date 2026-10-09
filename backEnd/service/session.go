package service

import (
	"backend/model"
	"backend/repos"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func CreateSession(db *sql.DB, userID int64) (string, error) {

	sessionID := uuid.New().String()
	expiredAt := time.Now().Add(24 * time.Hour)

	var session model.Session

	session.UserID = int(userID)
	session.SessionID = sessionID
	session.ExpiredAt = expiredAt

	if err := repos.CreateSession(db, session); err != nil {
		return "", err
	}

	return sessionID, nil
}
