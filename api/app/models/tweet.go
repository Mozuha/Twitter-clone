package models

import "time"

type Tweet struct {
	Id        uint      `json:"id"`
	ParentId  uint      `json:"parent_id"`
	UserId    uint      `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
