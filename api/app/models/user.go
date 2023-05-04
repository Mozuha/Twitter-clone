package models

import "time"

type User struct {
	Id           uint      `json:"id"`
	DisplayName  string    `json:"display_name"`
	Username     string    `json:"username" binding:"required"`
	Email        string    `json:"email" binding:"required,email"`
	Password     string    `json:"password" binding:"required"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
