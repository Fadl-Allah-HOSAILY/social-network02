package model

import "time"

type User struct {
    UserID            int       `json:"userId"`
    FirstName         string    `json:"firstName"`
    LastName          string    `json:"lastName"`
    Nickname          string    `json:"nickname"`
    Avatar            string    `json:"avatar"`
    AboutMe           string    `json:"aboutMe"`
    ProfileVisibility string    `json:"profileVisibility"`
    Email             string    `json:"email"`
    Password          string    `json:"password"`
    Birthday          time.Time `json:"birthday"`
    CreatedAt         time.Time `json:"createdAt"`
}