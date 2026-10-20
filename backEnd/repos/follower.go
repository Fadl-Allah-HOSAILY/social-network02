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
