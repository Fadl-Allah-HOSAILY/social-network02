package model

import "time"

type RegisterRequest  struct {
    FirstName         string    `json:"firstName"`
    LastName          string    `json:"lastName"`
    Nickname          string    `json:"nickname"`
    Avatar            string    `json:"avatar"`
    AboutMe           string    `json:"aboutMe"`
    Email             string    `json:"email"`
    Password          string    `json:"password"`
    Birthday          time.Time `json:"birthday"`
}