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
