package model

import "time"

type ProfileResponse struct {
    UserID            int       `json:"userId"`
    FirstName         string    `json:"firstName"`
    LastName          string    `json:"lastName"`
    Nickname          string    `json:"nickname"`
    Avatar            string    `json:"avatar"`
    AboutMe           string    `json:"aboutMe"`
    ProfileVisibility string    `json:"profileVisibility"`
    Birthday          time.Time `json:"birthday"`
}