package service

import (
	"backend/model"
	"errors"
	"regexp"
	"strings"
	"time"
)

func Register(userData model.RegisterRequest) error {

	letterRegex := regexp.MustCompile(`[^\p{L}]`)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	passwordRegex := regexp.MustCompile(`[!@#$%^&*]`)
	numberRegex := regexp.MustCompile(`[0-9]`)

	firstName := strings.TrimSpace(userData.FirstName)

	if len(firstName) < 3 {
		return errors.New("First Name should be more than 2 character")
	} else if len(firstName) > 30 {
		return errors.New("First Name should be less than 30 character")
	} else if letterRegex.MatchString(firstName) {
		return errors.New("Special Character or Numbers Not Allowed In First Name here")
	}

	lastName := strings.TrimSpace(userData.LastName)

	if len(lastName) < 3 {
		return errors.New("Last Name should be more than 2 character")
	} else if len(lastName) > 30 {
		return errors.New("Last Name should be less than 30 character")
	} else if letterRegex.MatchString(lastName) {
		return errors.New("Special Character or Numbers Not Allowed In Last Name here")
	}

	email := strings.TrimSpace(userData.Email)

	if !emailRegex.MatchString(email) {
		return errors.New("Invalid email format")
	}

	password := userData.Password

	if len(password) < 8 {
		return errors.New("Password should be at least 8 characters")
	} else if len(password) > 30 {
		return errors.New("Password should be less than 30 characters")
	} else if !passwordRegex.MatchString(password) {
		return errors.New("Special Character Required")
	} else if !numberRegex.MatchString(password) {
		return errors.New("Numbers Required")
	}

	if userData.Birthday.IsZero() || userData.Birthday.After(time.Now()) {
		return errors.New("Birthday Not Valid")
	}

	nickName := strings.TrimSpace(userData.Nickname)

	if nickName != "" {
		if len(nickName) < 3 {
			return errors.New("NickName should be at least 3 characters")
		} else if len(nickName) > 30 {
			return errors.New("NickName should be less than 30 characters")
		}
	}

	aboutMe := strings.TrimSpace(userData.AboutMe)

	if len(aboutMe) > 500 {
		return errors.New("aboutMe should be less than 500 characters")
	}

	return nil
}