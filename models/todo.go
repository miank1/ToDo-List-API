package models

import "time"

type Todo struct {
	ID          uint      `gorm:"primaryKey" json:"id"` // Primary key
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // pending / completed
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
