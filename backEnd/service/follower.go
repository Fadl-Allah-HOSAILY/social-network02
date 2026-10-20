
package service

import (
	"backend/model"
	"backend/repos"
	"database/sql"
	"errors"
)

var ErrCannotFollowSelf = errors.New("cannot follow yourself")
var ErrFollowerNotFound = errors.New("cannot Find Follower")
var ErrFollowingNotFound = errors.New("cannot Find Following")

func FollowUser(db *sql.DB, followerID int, followingID int) (string, error) {
	if followerID == followingID {
		return "", ErrCannotFollowSelf
	}

	var follower model.User
	var following model.User

	follower, err := repos.GetUserByID(db, followerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrFollowerNotFound
		}
		return "", err
	}

	following, err = repos.GetUserByID(db, followingID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrFollowingNotFound
		}
		return "", err
	}

	follows, err := repos.FollowExists(db, followerID, followingID)
	if err != nil {
		return "", err
	}

	if follows {
		err = repos.Unfollow(db, followerID, followingID)
		if err != nil {
			return "", err
		}
		return "unfollowed", nil
	}

	status := "pending"
	if following.ProfileVisibility == "public" {
		status = "accepted"
	}

	err = repos.CreateFollow(db, follower.UserID, following.UserID, status)
	if err != nil {
		return "", err
	}

	return "followed", nil
}
