package repos

import (
	"database/sql"
	"errors"
)

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

func FollowExists(db *sql.DB, followerID int, followingID int) (bool, error) {
	row := db.QueryRow(`
	SELECT EXISTS(
    SELECT 1
    FROM followers
    WHERE follower_id = ? AND following_id = ?
	);
	`, followerID, followingID)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

var ErrFollowNotFound = errors.New("follow relationship not found")

func DeleteFollow(db *sql.DB, followerID int, followingID int) error {
	res, err := db.Exec(`
	DELETE FROM followers
	WHERE follower_id = ? AND following_id = ?
	`, followerID, followingID)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrFollowNotFound
	}

	return nil
}
