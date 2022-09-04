package model

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	// TODO: Add more fields
}

func (User) TableName() string { return "users" }
