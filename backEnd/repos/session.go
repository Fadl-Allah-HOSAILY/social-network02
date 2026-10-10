package repos

import (
	"backend/model"
	"database/sql"
)

func CreateSession(db *sql.DB, session model.Session) error {
	_, err := db.Exec(`
		INSERT INTO Sessions (User_id, session_id, expired_at)
		VALUES(?,?,?)
	`, session.UserID, session.SessionID, session.ExpiredAt)

	if err != nil {
		return err
	}

	return nil
}

func GetSession(db *sql.DB, sessionID string) (model.Session, error) {
	var session model.Session

	row := db.QueryRow(`
		SELECT user_id, session_id, expired_at
		FROM Sessions
		WHERE session_id = ?;
	`, sessionID)

	err := row.Scan(&session.UserID, &session.SessionID, &session.ExpiredAt)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}
