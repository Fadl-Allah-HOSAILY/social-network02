package service

import (
	"backend/model"
	"backend/repos"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrSessionExpired = errors.New("session expired")
var ErrSessionNotFound = errors.New("session not found")

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

func CheckSession(db *sql.DB, sessionID string) (model.Session, error) {
	session, err := repos.GetSession(db, sessionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Session{}, ErrSessionNotFound
		}

		return model.Session{}, err
	}

	if !session.ExpiredAt.After(time.Now()) {
		return model.Session{}, ErrSessionExpired

	}

	return session, nil
}
