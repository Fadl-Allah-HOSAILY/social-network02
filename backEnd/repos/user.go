package repos

import (
	"database/sql"
	"backend/model"
)

func CreateUser(db *sql.DB, user model.RegisterRequest) error{

	_, err := db.Exec(`
		INSERT INTO Users (first_name, last_name, nickname, avatar, about_me, email, password, birthday)
		VALUES(?,?,?,?,?,?,?,?)
	`, user.FirstName, user.LastName, user.Nickname, user.Avatar, user.AboutMe, user.Email, user.Password, user.Birthday.Format("2006-01-02"))

	if err != nil {
		return err
	}

	return nil
}