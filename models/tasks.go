package models

import "time"

type Task struct {
	ID          uint16    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	User_id     uint16    `json:"user_id" gorm:"foreignKey:ID"`
}

func (Task) TableName() string {
	return "tasks"
}
