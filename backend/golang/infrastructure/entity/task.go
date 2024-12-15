package entity

import (
	"time"
)

type Task struct {
	TaskID    string `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
