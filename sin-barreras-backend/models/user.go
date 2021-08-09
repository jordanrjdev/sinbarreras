package models

import "time"

type User struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Avatar    string     `json:"avatar"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
