package service

import (
	"backend/model"
	"backend/repos"
	"database/sql"
	"errors"
)

var ErrPrivateProfile = errors.New("private profile")
var ErrProfileNotFound = errors.New("profile not found")

func GetProfile(db *sql.DB, profileID int, viewerID int) (model.User, error) {
	user, err := repos.GetUserByID(db, profileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, ErrProfileNotFound
		}
		return model.User{}, err
	}

	if user.ProfileVisibility != "public" && viewerID != profileID {
		return model.User{}, ErrPrivateProfile
	}

	return user, nil
}
