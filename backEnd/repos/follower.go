package repos

import "database/sql"

func CreateFollow(db *sql.DB, followerID int, followingID int, status string) error {

	_, err := db.Exec(`
	INSERT INTO followers (follower_id, following_id, status)
	VALUES(?,?,?)
	`, followerID, followingID, status)

	if err != nil {
		return err
	}

	return nil
}

