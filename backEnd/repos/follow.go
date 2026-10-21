package repos

import (
	"backend/model"
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

func Unfollow(db *sql.DB, followerID int, followingID int) error {
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

func FetchFollow(db *sql.DB, action string, userID int) ([]model.User, error) {
	var users []model.User
	switch action {

	case "followers":
		rows, err := db.Query(`
		SELECT 
		f.follower_id,
		u.first_name,
		u.last_name,
		u.nickname,
		u.avatar
		FROM followers f
		JOIN Users u ON f.follower_id = u.user_id
		WHERE f.following_id = ? 
		AND f.status = 'accepted'  
		`, userID)

		if err != nil {
			return []model.User{}, err
		}

		defer rows.Close()

		for rows.Next() {
			var user model.User
			err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Nickname, &user.Avatar)
			if err != nil {
				return []model.User{}, err
			}
			users = append(users, user)
		}
		err = rows.Err()
		if err != nil {
			return []model.User{}, err
		}

	case "following":
		rows, err := db.Query(`
		SELECT 
		f.following_id,
		u.first_name,
		u.last_name,
		u.nickname,
		u.avatar
		FROM followers f
		JOIN Users u ON f.following_id = u.user_id
		WHERE f.follower_id = ? 
		AND f.status = 'accepted'  
		`, userID)

		if err != nil {
			return []model.User{}, err
		}

		defer rows.Close()

		for rows.Next() {
			var user model.User
			err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Nickname, &user.Avatar)
			if err != nil {
				return []model.User{}, err
			}
			users = append(users, user)
		}
		err = rows.Err()
		if err != nil {
			return []model.User{}, err
		}

	default:
		return []model.User{}, errors.New("invalid action")
	}

	return users, nil
}

func UpdateFollowStatus(db *sql.DB, followerID int, followingID int, status string) error {
	res, err := db.Exec(`
	UPDATE Followers
	SET status = ?
	WHERE follower_id = ?
	AND following_id = ?
	AND status = 'pending';
	`,status, followerID, followingID)

	if err != nil {
		return err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsUpdated == 0 {
    return ErrFollowNotFound
}
	
	return nil 
}
